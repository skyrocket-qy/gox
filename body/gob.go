package body

import (
	"bytes"
	"encoding/gob"

	"github.com/skyrocket-qy/erx"
)

func GobEncode(in any) ([]byte, error) {
	var buf bytes.Buffer

	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(in); err != nil {
		return nil, erx.W(err)
	}

	return buf.Bytes(), nil
}

func GobDecode(src []byte, target any) error {
	dec := gob.NewDecoder(bytes.NewReader(src))
	if err := dec.Decode(target); err != nil {
		return erx.W(err)
	}

	return nil
}
