package drivers

import "errors"

var ErrInvalidConfigStruct = errors.New("invalid config structure")
var ErrEmptyStruct = errors.New("no empty struct")
var ErrNotModified          = errors.New("resource was not modified")