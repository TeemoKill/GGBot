package consts

import (
	"errors"
	"fmt"
)

func GGError(format string, args ...interface{}) (err error) {
	err = errors.New(fmt.Sprintf(format, args...))
	return err
}

var (
	ErrorNotImplemented = GGError("not implemented")
)
