package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"streamingservice/pkg/model"

	"github.com/nats-io/stan.go"
)

func (s *Server) natsConn(clientID, clusterID, natsurl string) error {
	sc, err := stan.Connect(clusterID, clientID, stan.NatsURL(natsurl),
		stan.Pings(1, 3),
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			log.Fatalf("Connection lost, reason: %v", reason)
			s.Store.Order().GetAll(context.Background())
		}))
	if err != nil {
		return err
	}
	fmt.Println("Nats connected")

	s.stanConn = sc

	return nil
}

func (s *Server) GetMessages(subject, qgroup, durable string) error {
	msgHandler := func(msg *stan.Msg) {
		if err := msg.Ack(); err != nil {
			log.Printf("failed to ACK msg:%v", err)
			return
		}

		var order model.Order
		err := json.Unmarshal(msg.Data, &order)
		if err != nil {
			log.Println("Fail to decode message:", err)
			return
		}
		order.Delivery.OrderID = order.OrderUID
		order.Payment.OrderID = order.OrderUID
		for i := range order.Items {
			order.Items[i].OrderID = order.Delivery.OrderID
		}

		if err := order.Validate(); err != nil {
			log.Println(err)
			return
		}

		if err := s.Store.Order().Create(context.Background(), &order); err != nil {
			log.Println(err)
		}
	}

	_, err := s.stanConn.QueueSubscribe(subject,
		qgroup, msgHandler,
		stan.DeliverAllAvailable(),
		stan.SetManualAckMode(),
		stan.DurableName(durable))
	if err != nil {
		log.Println(err)
	}
	return nil
}
