package cache

import "streamingservice/pkg/store"

type Cache struct {
	Data            map[string]interface{}
	store           store.Store
	orderRepository *OrderRepository
}

func NewCache(store store.Store) *Cache {
	return &Cache{
		Data:  make(map[string]interface{}),
		store: store,
	}
}

func (c *Cache) Order() store.OrderRepository {
	if c.orderRepository != nil {
		return c.orderRepository
	}

	c.orderRepository = &OrderRepository{
		cache: c,
	}

	return c.orderRepository
}
