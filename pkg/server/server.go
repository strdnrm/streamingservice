package server

import (
	"context"
	"fmt"
	"streamingservice/pkg/config"
	"streamingservice/pkg/routes"
	"streamingservice/pkg/store"
	"streamingservice/pkg/store/cache"
	"streamingservice/pkg/store/sqlstore"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/nats-io/stan.go"
)

type Server struct {
	config   config.Config
	stanConn stan.Conn
	Router   *gin.Engine
	Store    store.Store
	Cache    cache.Cache
}

func New(c config.Config) (*Server, error) {
	db, err := newDB(c.Dburl)
	if err != nil {
		return nil, err
	}
	fmt.Println("Postgres connected")
	// defer db.Close()

	return &Server{
		config: c,
		Store:  cache.NewCache(sqlstore.New(db)),
		Router: routes.ConfigureRouter(cache.NewCache(sqlstore.New(db))),
	}, nil
}

func (s *Server) Start() error {
	if err := s.natsConn(s.config.ClientID, s.config.ClusterID, s.config.Natsurl); err != nil {
		return err
	}

	_, err := s.Store.Order().GetAll(context.Background())
	if err != nil {
		return err
	}

	s.GetMessages("test", "test", "test-1")

	return s.Router.Run(s.config.Port)
}

func newDB(databaseURL string) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		return nil, err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://./migrations",
		"postgres", driver,
	)
	if err != nil {
		return nil, err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return nil, err
	}

	return db, nil
}
