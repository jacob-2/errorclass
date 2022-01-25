package errorclass

import (
	"errors"
	"fmt"
	"testing"
)

var err error = ErrEndUser{ErrDoesntExist{fmt.Errorf("wrapped error")}}
func TestErrorsDotIs(t *testing.T) {
	if !errors.Is(err, ErrEndUser{}) {
		t.Error("outermost error failed to match for errors.Is")
	}
	if !errors.Is(err, ErrDoesntExist{}) {
		t.Error("inner error failed to match for errors.Is")
	}
	if errors.Is(err, ErrExists{}) {
		t.Error("errors.Is wrongly matched for other error type")
	}
}

func TestErrorsDotAs(t *testing.T) {
	err2 := ErrEndUser{}
	if !errors.As(err, &err2) {
		t.Error("expected errors.As to return true")
	}
	if err2.Error() != "ErrEndUser: ErrDoesntExist: wrapped error" {
		t.Log(err2.Error())
		t.Error("unwrapped error isn't correct")
	}
	err3 := ErrDoesntExist{}
	if !errors.As(err, &err3) {
		t.Error("expected errors.As to return true")
	}
	if err3.Error() != "ErrDoesntExist: wrapped error" {
		t.Log(err3.Error())
		t.Error("unwrapped error isn't correct")
	}
}

func TestErrorsDotUnwrap(t *testing.T) {
	if errors.Unwrap(err).Error() != "ErrDoesntExist: wrapped error" {
		t.Fail()
	}
}