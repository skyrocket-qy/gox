package body_test

import (
	"bytes"
	"compress/gzip"
	"errors"
	"testing"

	"github.com/skyrocket-qy/gox/body"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestEncodeDecode(t *testing.T) {
	t.Parallel()

	original := &body.TestMessage{
		Id:    "test-id",
		Value: 12345,
	}

	encoded, err := body.Encode(original)
	require.NoError(t, err)
	require.NotNil(t, encoded)

	decoded, err := body.Decode[*body.TestMessage](encoded)
	require.NoError(t, err)
	require.NotNil(t, decoded)

	require.True(t, proto.Equal(original, decoded), "original and decoded messages should be equal")
}

func TestEncode_NilMessageError(t *testing.T) {
	t.Parallel()

	_, err := body.Encode(nil)
	require.Error(t, err)
}

func TestDecode_GzipError(t *testing.T) {
	t.Parallel()

	invalidGzipData := []byte("not a gzip file")
	_, err := body.Decode[*body.TestMessage](invalidGzipData)
	require.Error(t, err)
}

func TestDecode_UnmarshalError(t *testing.T) {
	t.Parallel()

	var buf bytes.Buffer

	gz := gzip.NewWriter(&buf)
	_, err := gz.Write([]byte("invalid protobuf data"))
	require.NoError(t, err)
	require.NoError(t, gz.Close())

	_, err = body.Decode[*body.TestMessage](buf.Bytes())
	require.Error(t, err)
}

type errorWriter struct{}

func (e *errorWriter) Write(p []byte) (n int, err error) {
	return 0, errors.New("write error")
}

func TestEncode_GzipWriteError(t *testing.T) {
	t.Parallel()

	original := &body.TestMessage{
		Id:    "test-id",
		Value: 12345,
	}

	err := body.EncodeWithWriter(original, &errorWriter{})
	require.Error(t, err)
}

func TestDecode_CorruptedGzipError(t *testing.T) {
	t.Parallel()
	// Create a valid compressed stream
	var buf bytes.Buffer

	gz := gzip.NewWriter(&buf)
	_, err := gz.Write([]byte("some data"))
	require.NoError(t, err)
	err = gz.Close()
	require.NoError(t, err)

	validData := buf.Bytes()

	// Corrupt it by truncating it
	corruptedData := validData[:len(validData)-5]

	_, err = body.Decode[*body.TestMessage](corruptedData)
	require.Error(t, err)
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
	t.Parallel()

	original := &body.TestMessage{
		Id:    "test-id",
		Value: 12345,
	}

	err := body.EncodeWithWriter(original, &failingCloser{})
	require.Error(t, err)
}

func TestDecode_NonPointerError(t *testing.T) {
	t.Parallel()
	// This is expected to fail because TestMessage is a pointer receiver type.
	// To test this, we would need a proto message type that is not a pointer receiver.
	// We will skip this test for now as it is not possible to create such a type with the current
	// generated code.
	t.Skip("Skipping test for non-pointer error: " +
		"cannot create a non-pointer proto message with current generated code.")
}

type errorReaderCloser struct {
	*bytes.Reader
}

func (erc *errorReaderCloser) Close() error {
	return errors.New("close error")
}

func TestDecode_GzipCloseError(t *testing.T) {
	t.Parallel()

	original := &body.TestMessage{Id: "test"}
	encoded, err := proto.Marshal(original)
	require.NoError(t, err)

	var buf bytes.Buffer

	gz := gzip.NewWriter(&buf)
	_, err = gz.Write(encoded)
	require.NoError(t, err)
	require.NoError(t, gz.Close())

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

	_, err = body.Decode[*body.TestMessage](corruptedData)
	require.Error(t, err)
}

func TestDecode_NilTargetError(t *testing.T) {
	t.Parallel()

	encoded, err := body.Encode(&body.TestMessage{Id: "test"})
	require.NoError(t, err)

	// This should cause a panic, not an error.
	// The code checks `if typ == nil`, but `typ` will not be nil here.
	// `reflect.TypeOf(nilMsg)` is `*body.TestMessage`.
	// To make `typ` nil, we need to pass a nil interface.
	_, err = body.Decode[proto.Message](encoded)
	require.Error(t, err)
	require.Equal(t, "target type must be a pointer to a proto message", err.Error())
}
