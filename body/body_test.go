package body_test

import (
	"testing"

	"github.com/skyrocket-qy/gox/body"
	pkgpbv1 "github.com/skyrocket-qy/protos/gen/pkgpb/v1"
)

func TestEncode(t *testing.T) {
	in := pkgpbv1.CursorData{}
	if _, err := body.Encode(&in); err != nil {
		t.Fatalf("Encode failed: %v", err)
	}
}

func TestDecode(t *testing.T) {
	var in []byte
	if _, err := body.Decode[*pkgpbv1.CursorData](in); err != nil {
		t.Fatalf("Decode failed: %v", err)
	}
}
