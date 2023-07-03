package model

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Order struct {
	OrderUID          string    `db:"order_uid" json:"order_uid" validate:"required,min=19,max=19"`
	TrackNumber       string    `db:"track_number" json:"track_number" validate:"required"`
	Entry             string    `db:"entry" json:"entry" validate:"required"`
	Delivery          Delivery  `json:"delivery" validate:"required"`
	Payment           Payment   `json:"payment" validate:"required"`
	Items             []Item    `json:"items" validate:"required"`
	Locale            string    `db:"locale" json:"locale" validate:"required"`
	InternalSignature string    `db:"internal_signature" json:"internal_signature"`
	CustomerID        string    `db:"customer_id" json:"customer_id" validate:"required"`
	DeliveryService   string    `db:"delivery_service" json:"delivery_service" validate:"required"`
	ShardKey          string    `db:"shardkey" json:"shardkey" validate:"required"`
	SMID              int       `db:"sm_id" json:"sm_id" validate:"required"`
	DateCreated       time.Time `db:"date_created" json:"date_created" validate:"required"`
	OOFShard          string    `db:"oof_shard" json:"oof_shard" validate:"required"`
}

func (o *Order) Validate() error {
	validate := validator.New()
	err := validate.Struct(o.Delivery)
	if err != nil {
		return err
	}
	err = validate.Struct(o.Payment)
	if err != nil {
		return err
	}
	for i := range o.Items {
		err := validate.Struct(o.Items[i])
		if err != nil {
			return err
		}
	}
	err = validate.Struct(o)
	if err != nil {
		return err
	}
	return nil
}
