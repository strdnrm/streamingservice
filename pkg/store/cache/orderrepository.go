package cache

import (
	"context"
	"errors"
	"streamingservice/pkg/model"
)

type OrderRepository struct {
	cache *Cache
}

func (r *OrderRepository) Create(ctx context.Context, order *model.Order) error {
	r.cache.Data[order.OrderUID] = order
	if err := r.cache.store.Order().Create(ctx, order); err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) GetById(ctx context.Context, id string) (*model.Order, error) {
	data, found := r.cache.Data[id]
	if found {
		order, ok := data.(model.Order)
		if !ok {
			return nil, errors.New("Corrupted data")
		}
		return &order, nil
	} else {
		order, err := r.cache.store.Order().GetById(ctx, id)
		if err != nil {
			return nil, err
		}
		r.cache.Data[order.OrderUID] = order
		return order, nil
	}
}

func (r *OrderRepository) GetAll(ctx context.Context) (*[]model.Order, error) {
	orders, err := r.cache.store.Order().GetAll(context.Background())
	if err != nil {
		return nil, err
	}
	for _, order := range *orders {
		r.cache.Data[order.OrderUID] = order
	}
	return orders, nil
}
