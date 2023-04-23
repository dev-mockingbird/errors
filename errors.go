package errors

import (
	"errors"
	"fmt"
)

type noticeError struct {
	msg string
}

func (n noticeError) Error() string {
	return n.msg
}

func IsNotice(err error) bool {
	_, ok := err.(noticeError)
	return ok
}

func Notice(msg string) error {
	return noticeError{msg: msg}
}

func Noticef(format string, v ...any) error {
	return noticeError{msg: fmt.Sprintf(format, v...)}
}

func Wrap(err error, msg string) error {
	return fmt.Errorf("%s: %w", msg, err)
}

func New(msg string) error {
	return errors.New(msg)
}

func Is(err, target error) bool {
	return errors.Is(err, target)
}

func As(err error, target any) bool {
	return errors.As(err, target)
}
