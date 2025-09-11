package lib

import "testing"

func TestAs(t *testing.T) {
	if As("foo", "bar") != "foo AS bar" {
		t.Error("As function failed")
	}
}

func TestOrd(t *testing.T) {
	if Ord("foo", true) != "foo" {
		t.Error("Ord with true failed")
	}

	if Ord("foo", false) != "foo DESC" {
		t.Error("Ord with false failed")
	}
}
