package data

import (
	"github.com/go-kratos/kratos/v2/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	"nacos/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo, NewDB)

// Data .
type Data struct {
	// TODO wrapped database client
	db  *sqlx.DB
	log *log.Helper
}

// NewDB .
func NewDB(c *conf.Data, logger log.Logger) *sqlx.DB {
	log := log.NewHelper(log.With(logger, "module", "data/sqlx"))

	db, err := sqlx.Connect(c.Database.Driver, c.Database.Source)
	if err != nil {
		log.Fatal("create sqlx.DB failed, err:", err)
	}

	return db
}

// NewData .
func NewData(db *sqlx.DB, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(log.With(logger, "module", "data"))

	cleanup := func() {
		log.Info("closing the data resources")
	}

	return &Data{
		db:  db,
		log: log,
	}, cleanup, nil
}
