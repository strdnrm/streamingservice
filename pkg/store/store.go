package store

type Store interface {
	Order() OrderRepository
}

// type Cache interface {
// 	Order() OrderRepository
// }
