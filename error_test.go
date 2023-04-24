package errors

import (
	"testing"
)

func TestPrefixedError(t *testing.T) {
	err := New("hello world", InternalError)
	err = Wrap(Wrap(Wrap(err, "wrapped"), "wrapped 1"), "wrapped 2")
	msg := err.Error()
	if msg != "wrapped 2: wrapped 1: wrapped: [internal] hello world" {
		t.Fatal("Error")
	}
	prefix, msg := Parse(err)
	if prefix != InternalError || msg != "hello world" {
		t.Fatal("unwrap prefix")
	}
}
