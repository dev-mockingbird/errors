package errors

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

const (
	InternalError = "internal"
	PrefixNone    = ""
)

type prefixedError struct {
	code string
	msg  string
}

func (e prefixedError) Error() string {
	if e.code == "" {
		return e.msg
	}
	return fmt.Sprintf("[%s] %s", e.code, e.msg)
}

func Ancestor(err error) error {
	origin := err.Error()
	if idx := strings.LastIndex(origin, ": "); idx > -1 && idx < len(origin)-2 {
		origin = origin[idx+2:]
		return errors.New(origin)
	}
	return err
}

func Parse(err error) (code string, msg string) {
	origin := Ancestor(err).Error()
	reg := regexp.MustCompile(`^\[.+\] `)
	msg = reg.ReplaceAllStringFunc(origin, func(matched string) string {
		code = regexp.MustCompile(`[\[\]]`).ReplaceAllString(strings.Trim(matched, " "), "")
		return ""
	})
	return
}

func Wrap(err error, msg string) error {
	return fmt.Errorf("%s: %w", msg, err)
}

func Unwrap(err error) error {
	msg := err.Error()
	if idx := strings.Index(msg, ": "); idx > -1 && idx < len(msg)-2 {
		return New(msg[idx+2:])
	}
	return err
}

func New(msg string, code ...string) error {
	return prefixedError{msg: msg, code: func() string {
		if len(code) > 0 {
			return code[0]
		}
		return ""
	}()}
}

func Is(err, target error) bool {
	return errors.Is(err, target)
}

func As(err error, target any) bool {
	return errors.As(err, target)
}
