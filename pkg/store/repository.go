package store

import (
	"context"
	"streamingservice/pkg/model"
)

type OrderRepository interface {
	Create(ctx context.Context, order *model.Order) error
	GetById(ctx context.Context, id string) (*model.Order, error)
	GetAll(ctx context.Context) (*[]model.Order, error)
}

// type StorageRepository interface {
// 	CreateOrder(ctx context.Context, order *model.Order) error
// 	GetOrderById(ctx context.Context, id string) (*model.Order, error)
// 	GetOrders(ctx context.Context) (*[]model.Order, error)
// }
