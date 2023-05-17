package errors

import "strings"

type JoinableError interface {
	Join(err error)
	Error() string
}

type multiError []error

var _ error = multiError{}

func MultiError() JoinableError {
	return &multiError{}
}

func (me *multiError) Join(err error) {
	*me = append(*me, err)
}

func (me multiError) Error() string {
	errs := make([]string, len(me))
	for i, e := range me {
		errs[i] = e.Error()
	}
	return strings.Join(errs, "\n")
}
