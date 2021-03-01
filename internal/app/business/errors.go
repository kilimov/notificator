package business

import "errors"

var (
	UserCanNotBeEmpty = errors.New("user can not be empty")
	InvalidId         = errors.New("id nit exist")
)
