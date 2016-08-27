package errors

import (
	stderrs "errors"
)

var (
	NotImplemented = stderrs.New("Not implemented")
)
