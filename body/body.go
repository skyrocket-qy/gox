package body

import (
	"bytes"
	"compress/gzip"
	"errors"
	"io"

	"google.golang.org/protobuf/proto"
)

func Encode(in proto.Message) ([]byte, error) {
	raw, err := proto.Marshal(in)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer

	gz := gzip.NewWriter(&buf)

	if _, err := gz.Write(raw); err != nil {
		return nil, err
	}

	if err := gz.Close(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func Decode[T proto.Message](data []byte) (T, error) {
	var out T

	buf := bytes.NewReader(data)

	gz, err := gzip.NewReader(buf)
	if err != nil {
		return out, err
	}

	defer func() {
		if cerr := gz.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	decompressed, err := io.ReadAll(gz)
	if err != nil {
		return out, err
	}

	msg := proto.Clone(out)
	if err := proto.Unmarshal(decompressed, msg); err != nil {
		return out, err
	}

	if assertedMsg, ok := msg.(T); ok {
		return assertedMsg, nil
	}

	return out, errors.New("failed to assert type of decoded message")
}
