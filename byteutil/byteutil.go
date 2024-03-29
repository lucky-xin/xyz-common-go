package byteutil

import (
	"bytes"
	"encoding/binary"
)

func ReadBytes(payload []byte, start int) (int, []byte) {
	l := int(binary.BigEndian.Uint32(payload[start : start+4]))
	byts := payload[start+4 : start+4+l]
	return l, byts
}

func WriteBytes(buffer *bytes.Buffer, byes []byte) (err error) {
	lenBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(lenBytes, uint32(len(byes)))
	_, err = buffer.Write(lenBytes)
	if err != nil {
		return
	}
	_, err = buffer.Write(byes)
	return
}
