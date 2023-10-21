package errors

import (
	"errors"
	"fmt"
	"strings"

	"github.com/thoas/go-funk"
)

type HasCode interface {
	GetCode() string
}

type taggedError struct {
	tags []string
	err  error
}

func (e taggedError) GetCode() string {
	return strings.Join(e.tags, ";")
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

func Tags(err error) []string {
	tags, ok := err.(interface {
		Tags() []string
	})
	if ok {
		return tags.Tags()
	}
	return nil
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
			ret = e
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
			ret = e
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
