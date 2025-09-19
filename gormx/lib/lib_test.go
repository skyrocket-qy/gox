package lib_test

import (
	"testing"

	"github.com/skyrocket-qy/gox/gormx/lib"
)

func TestAs(t *testing.T) {
	if lib.As("foo", "bar") != "foo AS bar" {
		t.Error("As function failed")
	}
}

func TestOrd(t *testing.T) {
	if lib.Ord("foo", true) != "foo" {
		t.Error("Ord with true failed")
	}

	if lib.Ord("foo", false) != "foo DESC" {
		t.Error("Ord with false failed")
	}
}
