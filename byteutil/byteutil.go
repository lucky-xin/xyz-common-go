package byteutil

import (
	"bytes"
	"encoding/binary"
)

func ReadBytes(payload []byte, start uint32) (uint32, []byte) {
	uuidLen := binary.BigEndian.Uint32(payload[start : start+4])
	uuidBytes := payload[start+4 : start+4+uuidLen]
	return uuidLen, uuidBytes
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
