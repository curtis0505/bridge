package types

import (
	"errors"
	"fmt"
)

var (
	NotImplemented     = errors.New("not implemented")
	NotSupported       = errors.New("not supported")
	InvalidTransaction = errors.New("invalid transaction")
	InvalidParameter   = errors.New("invalid parameter")
)

func WrapError(msg string, err error) error {
	return fmt.Errorf("%s: %s", msg, err.Error())
}
