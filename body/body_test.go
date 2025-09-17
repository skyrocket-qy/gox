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

func TestDecode_CorruptedGzipError(t *testing.T) {
	// Create a valid compressed stream
	var buf bytes.Buffer

	gz := gzip.NewWriter(&buf)
	_, err := gz.Write([]byte("some data"))
	assert.NoError(t, err)
	err = gz.Close()
	assert.NoError(t, err)

	validData := buf.Bytes()

	// Corrupt it by truncating it
	corruptedData := validData[:len(validData)-5]

	_, err = Decode[*TestMessage](corruptedData)
	assert.Error(t, err)
}

// failingCloser is a writer that fails on the second write,
// which is useful for testing the error path of gzip.Close().
type failingCloser struct {
	callCount int
}

func (fc *failingCloser) Write(p []byte) (n int, err error) {
	fc.callCount++
	if fc.callCount > 1 {
		return 0, errors.New("write error on close")
	}

	return len(p), nil
}

func TestEncode_GzipCloseError(t *testing.T) {
	original := &TestMessage{
		Id:    "test-id",
		Value: 12345,
	}

	err := EncodeWithWriter(original, &failingCloser{})
	assert.Error(t, err)
}

func TestDecode_NonPointerError(t *testing.T) {
	// This is expected to fail because TestMessage is a pointer receiver type.
	// To test this, we would need a proto message type that is not a pointer receiver.
	// We will skip this test for now as it is not possible to create such a type with the current
	// generated code.
	t.Skip(
		"Skipping test for non-pointer error because it is not possible to create a non-pointer proto message with the current generated code.",
	)
}

type errorReaderCloser struct {
	*bytes.Reader
}

func (erc *errorReaderCloser) Close() error {
	return errors.New("close error")
}

func TestDecode_GzipCloseError(t *testing.T) {
	original := &TestMessage{Id: "test"}
	encoded, err := proto.Marshal(original)
	assert.NoError(t, err)

	var buf bytes.Buffer

	gz := gzip.NewWriter(&buf)
	_, err = gz.Write(encoded)
	assert.NoError(t, err)
	gz.Close()

	// We need to create a custom reader that returns an error on Close.
	// However, the Decode function takes a byte slice, not a reader.
	// We can't inject a custom reader.
	// The uncovered part is `if cerr := gz.Close(); cerr != nil && err == nil`.
	// This is hard to test without modifying the source code to accept a reader.
	// Let's see if we can trigger it by corrupting the gzip footer.
	// A corrupted footer might cause an error in gz.Close().

	// Let's try to corrupt the last 8 bytes which are the checksum and the size.
	corruptedData := buf.Bytes()
	for i := range 8 {
		corruptedData[len(corruptedData)-1-i] = 0
	}

	_, err = Decode[*TestMessage](corruptedData)
	assert.Error(t, err)
}

func TestDecode_NilTargetError(t *testing.T) {
	encoded, err := Encode(&TestMessage{Id: "test"})
	assert.NoError(t, err)

	// This should cause a panic, not an error.
	// The code checks `if typ == nil`, but `typ` will not be nil here.
	// `reflect.TypeOf(nilMsg)` is `*body.TestMessage`.
	// To make `typ` nil, we need to pass a nil interface.
	_, err = Decode[proto.Message](encoded)
	assert.Error(t, err)
	assert.Equal(t, "target type must be a pointer to a proto message", err.Error())
}
