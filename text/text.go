package text

import (
	"bytes"
	"encoding/binary"
	"errors"
)

type Segment struct {
	Length int
	Bytes  []byte
}

type Text struct {
	MagicByte *byte
	Segments  []*Segment
}

func FromBuffer(magicByte *byte, buffer *bytes.Buffer) (block *Text, err error) {
	var segments []*Segment
	payload := buffer.Bytes()
	start := 0
	if magicByte != nil {
		if payload[start] != *magicByte {
			err = errors.New("invalid magic byte")
			return
		}
		start++
	}

	for start < len(payload) {
		l, byts := ReadBytes(payload, start)
		segments = append(segments, &Segment{Bytes: byts, Length: l})
		start += l + 4
	}
	block = &Text{Segments: segments}
	return
}

func (block *Text) ToBuffer() (buff *bytes.Buffer, err error) {
	for i := range block.Segments {
		s := block.Segments[i]
		_, err = buff.Write(s.Bytes)
		if err != nil {
			return
		}
	}
	return
}

func ReadBytes(payload []byte, start int) (l int, byts []byte) {
	idx := start + 4
	l = int(binary.BigEndian.Uint32(payload[start:idx]))
	byts = payload[idx : idx+l]
	return l, byts
}
