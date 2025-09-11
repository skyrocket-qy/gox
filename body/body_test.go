package body

import (
	"bytes"
	"compress/gzip"
	"errors"
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

func TestEncode_NilMessageError(t *testing.T) {
	_, err := Encode(nil)
	assert.Error(t, err)
}

func TestDecode_GzipError(t *testing.T) {
	invalidGzipData := []byte("not a gzip file")
	_, err := Decode[*TestMessage](invalidGzipData)
	assert.Error(t, err)
}

func TestDecode_UnmarshalError(t *testing.T) {
	var buf bytes.Buffer

	gz := gzip.NewWriter(&buf)
	_, err := gz.Write([]byte("invalid protobuf data"))
	assert.NoError(t, err)
	gz.Close()

	_, err = Decode[*TestMessage](buf.Bytes())
	assert.Error(t, err)
}

type errorWriter struct{}

func (e *errorWriter) Write(p []byte) (n int, err error) {
	return 0, errors.New("write error")
}

func TestEncode_GzipWriteError(t *testing.T) {
	original := &TestMessage{
		Id:    "test-id",
		Value: 12345,
	}

	err := EncodeWithWriter(original, &errorWriter{})
	assert.Error(t, err)
}
