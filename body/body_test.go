package body

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
)

func TestEncodeDecode(t *testing.T) {
	original := &TestMessage{
		Id:    "test-id",
		Value: 12345,
	}

	encoded, err := Encode(original)
	assert.NoError(t, err)
	assert.NotNil(t, encoded)

	decoded, err := Decode[*TestMessage](encoded)
	assert.NoError(t, err)
	assert.NotNil(t, decoded)

	assert.True(t, proto.Equal(original, decoded), "original and decoded messages should be equal")
}
