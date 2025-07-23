package special

import "encoding/binary"

func AA() {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, 123456)
}
