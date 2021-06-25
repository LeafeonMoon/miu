package common

import (
	"errors"
)

func Error() error {
	return errors.New("error")
}
