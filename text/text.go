package text

import (
	"bytes"
	"errors"
	"github.com/lucky-xin/xyz-common-go/byteutil"
)

// Segment 文本段
type Segment struct {
	// Length 段长度
	Length int
	// Bytes 段内容
	Bytes []byte
}

// Block 文本数据块
type Block struct {
	// MagicByte 魔法标识
	MagicByte *byte
	// Segments 分段列表
	Segments []*Segment
}

func FromBuffer(magicByte *byte, buffer *bytes.Buffer) (block *Block, err error) {
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
		l, byts := byteutil.ReadBytes(payload, start)
		segments = append(segments, &Segment{Bytes: byts, Length: l})
		start += l + 4
	}
	block = &Block{Segments: segments}
	return
}

func (block *Block) ToBuffer() (buff bytes.Buffer, err error) {
	if block.MagicByte != nil {
		buff.WriteByte(*block.MagicByte)
	}
	for i := range block.Segments {
		s := block.Segments[i]
		err = byteutil.WriteBytes(&buff, s.Bytes)
		if err != nil {
			return
		}
	}
	return
}
