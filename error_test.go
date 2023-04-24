package errors

import (
	"testing"
)

func TestPrefixedError(t *testing.T) {
	err := New("hello world", InternalError)
	codes, msg := Parse(err)
	if !(len(codes) == 1 && codes[0] == InternalError && msg == "hello world") {
		t.Fatal("unwrap code")
	}
	err = Wrap(Wrap(Wrap(err, "wrapped"), "wrapped 1"), "wrapped 2")
	msg = err.Error()
	if msg != "wrapped 2: wrapped 1: wrapped: [internal] hello world" {
		t.Fatal("Error")
	}
	err = Tag(err, "code1", "code2")
	msg = err.Error()
	if msg != "[code1] [code2] wrapped 2: wrapped 1: wrapped: [internal] hello world" {
		t.Fatal("tag")
	}
	err = First(err, InternalError)
	if err.Error() != "[internal] hello world" {
		t.Fatal("first")
	}
}
