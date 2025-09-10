package body_test

import (
	"testing"

	"github.com/skyrocket-qy/gox/body"
	pkgpbv1 "github.com/skyrocket-qy/protos/gen/pkgpb/v1"
)

func TestEncode(t *testing.T) {
	in := pkgpbv1.CursorData{}
	body.Encode(&in)
}

func TestDecode(t *testing.T) {
	var in []byte
	body.Decode[*pkgpbv1.CursorData](in)
}
