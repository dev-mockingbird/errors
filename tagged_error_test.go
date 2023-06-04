// Copyright (c) 2023 Yang,Zhong
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package errors

import "testing"

func TestTaggedError(t *testing.T) {
	err := New("hello world", "test")
	if err.Error() != "[test] hello world" {
		t.Fatal("tag error")
	}
	err = Wrap(err, "wrap")
	if err.Error() != "wrap: [test] hello world" {
		t.Fatal("wrap error")
	}
	err = Tag(err, "wrap-tag")
	te := err.Error()
	if te != "[wrap-tag] wrap: [test] hello world" {
		t.Fatal("tag wrap error")
	}
	if err := FirstTagged(err, "wrap-tag"); err.Error() != "wrap: [test] hello world" {
		t.Fatal("first tagged error")
	}
	if err := LastTagged(err, "test"); err == nil || err != nil && err.Error() != "hello world" {
		t.Fatal("last tagged error")
	}
}
