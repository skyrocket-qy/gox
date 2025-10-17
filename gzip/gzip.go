package gzip

import (
	"bytes"
	"compress/gzip"
	"io"
)

func Encode(in []byte) (out []byte, err error) {
	buf := bytes.NewBuffer(nil)
	w := gzip.NewWriter(buf)

	if _, err = w.Write(in); err != nil {
		return nil, err
	}

	if err = w.Flush(); err != nil {
		return nil, err
	}

	if err = w.Close(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func Decode(in []byte) (out []byte, err error) {
	reader, err := gzip.NewReader(bytes.NewReader(in))
	if err != nil {
		return nil, err
	}

	out, err = io.ReadAll(reader)
	if err != nil {
		_ = reader.Close()

		return nil, err
	}

	if err = reader.Close(); err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return out, err
}
