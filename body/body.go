package body

import (
	"bytes"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"reflect"

	"google.golang.org/protobuf/proto"
)

func Encode(in proto.Message) ([]byte, error) {
	var buf bytes.Buffer

	err := EncodeWithWriter(in, &buf)

	return buf.Bytes(), err
}

func EncodeWithWriter(in proto.Message, w io.Writer) error {
	if in == nil {
		return errors.New("input message cannot be nil")
	}

	raw, err := proto.Marshal(in)
	if err != nil {
		return err
	}

	gz := gzip.NewWriter(w)

	if _, err := gz.Write(raw); err != nil {
		return err
	}

	if err := gz.Close(); err != nil {
		return err
	}

	return nil
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

	// We need a non-nil message to unmarshal into.
	// We create one using reflection since T is a generic type.
	typ := reflect.TypeOf(out)
	if typ == nil || typ.Kind() != reflect.Ptr {
		return out, errors.New("target type must be a pointer to a proto message")
	}

	msg, ok := reflect.New(typ.Elem()).Interface().(T)
	if !ok {
		return out, fmt.Errorf("failed to assert type %T to proto.Message", out)
	}

	if err := proto.Unmarshal(decompressed, msg); err != nil {
		return out, err
	}

	return msg, nil
}
