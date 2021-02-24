package database

import (
	"github.com/kilimov/notificator/internal/app/database/drivers"
	"github.com/kilimov/notificator/internal/app/database/drivers/mongo"
)

func New(conf drivers.DataStoreConfig) (drivers.DataStore, error) {
	if conf.DataStoreName == "mongo" {
		return mongo.New(conf)
	}

	return nil, ErrDatastoreNotImplemented
}
