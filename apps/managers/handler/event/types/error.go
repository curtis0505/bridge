package types

import "fmt"

func WrapError(event string, msg string, err error) error {
	return fmt.Errorf("%s: %s : %s", event, msg, err.Error())
}
