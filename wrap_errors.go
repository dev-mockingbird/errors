package errors

import (
	"errors"
	"fmt"
)

const (
	InternalError = "internal"
)

type wrapError struct {
	priv string
	from error
}

func (e wrapError) Error() string {
	err := ""
	if e.from != nil {
		err = ": " + e.from.Error()
	}
	return fmt.Sprintf("%s%s", e.priv, err)
}

func (e wrapError) Unwrap() error {
	return e.from
}

func Traverse(err error, do func(error) bool) {
	for {
		if !do(err) {
			return
		}
		unwrap, ok := err.(interface {
			Unwrap() error
		})
		if ok {
			err = unwrap.Unwrap()
			continue
		}
		break
	}
}

func Wrap(err error, msg string) error {
	return wrapError{from: err, priv: msg}
}

func Unwrap(err error) error {
	return errors.Unwrap(err)
}

func Is(err, target error) bool {
	return errors.Is(err, target)
}

func As(err error, target any) bool {
	return errors.As(err, target)
}
