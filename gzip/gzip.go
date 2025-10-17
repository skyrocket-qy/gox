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

	w.Flush()
	w.Close()

	return buf.Bytes(), nil
}

func Decode(in []byte) (out []byte, err error) {
	reader, err := gzip.NewReader(bytes.NewReader(in))
	if err != nil {
		return nil, err
	}

	out, err = io.ReadAll(reader)
	reader.Close()

	if err != nil {
		return nil, err
	}

	return out, err
}
