package errors

import (
	"strings"
	"sync"
)

type JoinableError interface {
	Join(err error)
	AsyncJoin(err error)
	Has() int
	Error() string
}

type multiError struct {
	errors []error
	lock   sync.RWMutex
}

var _ error = &multiError{}

func MultiError() JoinableError {
	return &multiError{}
}

func (me *multiError) Join(err error) {
	me.errors = append(me.errors, err)
}

func (me *multiError) AsyncJoin(err error) {
	me.lock.Lock()
	defer me.lock.Unlock()
	me.errors = append(me.errors, err)
}

func (me *multiError) Has() int {
	me.lock.RLock()
	defer me.lock.RUnlock()
	return len(me.errors)
}

func (me *multiError) Error() string {
	me.lock.RLock()
	defer me.lock.RUnlock()
	errs := make([]string, len(me.errors))
	for i, e := range me.errors {
		errs[i] = e.Error()
	}
	return strings.Join(errs, "\n")
}
