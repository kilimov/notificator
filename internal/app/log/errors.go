package log

import (
	stdlog "log"
)

type ErrorFn func() error

func ErrorFnPrintln(fn ErrorFn) {
	if err := fn(); err != nil {
		stdlog.Println(err)
	}
}
