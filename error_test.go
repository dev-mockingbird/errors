package errors

import (
	"errors"
	"testing"
)

func TestNotice(t *testing.T) {
	err := Notice("hello world")
	if !IsNotice(err) {
		t.Fatal("hello world")
	}
	err = errors.New("normal error")
	if IsNotice(err) {
		t.Fatal("normal error")
	}
}
