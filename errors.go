package errors

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

const (
	InternalError = "internal"
)

type taggedError struct {
	tags []string
	msg  string
}

func (e taggedError) Error() string {
	if len(e.tags) == 0 {
		return e.msg
	}
	en := make([]string, len(e.tags))
	for i, code := range e.tags {
		en[i] = "[" + code + "]"
	}

	return fmt.Sprintf("%s %s", strings.Join(en, " "), e.msg)
}

func Ancestor(err error) error {
	if err == nil {
		return nil
	}
	origin := err.Error()
	if idx := strings.LastIndex(origin, ": "); idx > -1 && idx < len(origin)-2 {
		origin = origin[idx+2:]
		return errors.New(origin)
	}
	return err
}

func Parse(err error) (codes []string, msg string) {
	if err == nil {
		return
	}
	origin := err.Error()
	var rest string
	if idx := strings.Index(origin, ": "); idx > -1 {
		origin = origin[:idx]
		rest = origin[idx:]
	}
	reg := regexp.MustCompile(`^\[.+\] `)
	msg = reg.ReplaceAllStringFunc(origin, func(matched string) string {
		codes = append(codes, regexp.MustCompile(`[\[\]]`).ReplaceAllString(strings.Trim(matched, " "), ""))
		return ""
	})
	msg += rest
	return
}

func Wrap(err error, msg string) error {
	return fmt.Errorf("%s: %w", msg, err)
}

func Unwrap(err error) error {
	if err == nil {
		return nil
	}
	msg := err.Error()
	if idx := strings.Index(msg, ": "); idx > -1 && idx < len(msg)-2 {
		return errors.New(msg[idx+2:])
	}
	return err
}

func Tag(err error, tag ...string) error {
	if err == nil {
		return nil
	}
	return New(err.Error(), tag...)
}

func FirstTagged(err error, tag string) error {
	if err == nil {
		return nil
	}
	msg := err.Error()
	if idx := strings.Index(msg, "["+tag+"]"); idx > -1 {
		if idx := strings.LastIndex(msg[:idx], ": "); idx > -1 {
			return errors.New(msg[idx+2:])
		}
		return err
	}
	return nil
}

func LastTagged(err error, tag string) error {
	if err == nil {
		return nil
	}
	msg := err.Error()
	if idx := strings.LastIndex(msg, "["+tag+"]"); idx > -1 {
		if idx := strings.LastIndex(msg[:idx], ": "); idx > -1 {
			return errors.New(msg[idx+2:])
		}
		return err
	}
	return nil
}

func New(msg string, tag ...string) error {
	return taggedError{msg: msg, tags: tag}
}

func Is(err, target error) bool {
	return errors.Is(err, target)
}

func As(err error, target any) bool {
	return errors.As(err, target)
}
