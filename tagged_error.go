package errors

import (
	"errors"
	"fmt"
	"strings"

	"github.com/thoas/go-funk"
)

type taggedError struct {
	tags []string
	err  error
}

func (e taggedError) Error() string {
	if len(e.tags) == 0 {
		return e.err.Error()
	}
	en := make([]string, len(e.tags))
	for i, code := range e.tags {
		en[i] = "[" + code + "]"
	}
	return fmt.Sprintf("%s %s", strings.Join(en, " "), e.err.Error())
}

func (e taggedError) Tags() []string {
	return e.tags
}

func (e taggedError) Unwrap() error {
	return e.err
}

func Parse(err error) (codes []string, msg string) {
	if taggedError, ok := err.(taggedError); ok {
		return taggedError.tags, taggedError.err.Error()
	}
	msg = err.Error()
	return
}

func Tag(err error, tag ...string) error {
	if err == nil {
		return nil
	}
	return taggedError{err: err, tags: tag}
}

func FirstTagged(err error, tag ...string) error {
	var ret error
	Traverse(err, func(err error) bool {
		if e, ok := err.(taggedError); ok {
			if len(tag) > 0 && !funk.ContainsString(e.tags, tag[0]) {
				return true
			}
			ret = e.err
			return false
		}
		return true
	})
	return ret
}

func LastTagged(err error, tag ...string) error {
	var ret error
	Traverse(err, func(err error) bool {
		if e, ok := err.(taggedError); ok {
			if len(tag) > 0 && !funk.ContainsString(e.tags, tag[0]) {
				return true
			}
			ret = e.err
		}
		return true
	})
	return ret
}

func New(msg string, tag ...string) error {
	err := errors.New(msg)
	if len(tag) == 0 {
		return err
	}
	return taggedError{err: err, tags: tag}
}
