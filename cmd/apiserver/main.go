package main

import (
	"github.com/kilimov/notificator/internal/app/apiserver"
)

var version = "unknown"

func main() {
	apiserver.Start(version)
}
