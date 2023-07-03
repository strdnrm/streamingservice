package sqlstore

import (
	"streamingservice/pkg/store"

	"github.com/jmoiron/sqlx"
)

type Store struct {
	db              *sqlx.DB
	orderRepository *OrderRepository
}

func New(db *sqlx.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) Order() store.OrderRepository {
	if s.orderRepository != nil {
		return s.orderRepository
	}

	s.orderRepository = &OrderRepository{
		store: s,
	}

	return s.orderRepository
}
