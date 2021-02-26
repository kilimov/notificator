package database

import (
	"errors"
)

var ErrDatastoreNotImplemented = errors.New("datastore not implemented")
var ErrEmptyStruct = errors.New("no empty struct")
