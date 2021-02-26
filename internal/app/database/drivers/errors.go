package drivers

import "errors"

var ErrInvalidConfigStruct = errors.New("invalid config structure")
var ErrUserDoesntExists = errors.New("user does not exists")
var ErrNotModified = errors.New("resource was not modified")
