package text

import "bytes"

type Segment struct {
	Begin int32
	End   int32
	Bytes []byte
}

type SegmentDef struct {
	Begin int32
	End   int32
}

type Text struct {
	Segments []*Segment
}

type Def struct {
	Defs []*SegmentDef
}

func FromBuffer(def *Def, buffer *bytes.Buffer) (block *Text, err error) {
	var segments []*Segment
	for i := range def.Defs {
		d := def.Defs[i]
		tmp := make([]byte, d.End-d.Begin+1)
		_, err = buffer.Read(tmp)
		if err != nil {
			return
		}
		segments = append(segments, &Segment{Begin: d.Begin, End: d.End, Bytes: tmp})
	}
	block = &Text{Segments: segments}
	return
}

func (block *Text) ToBuffer() (buff *bytes.Buffer) {
	for i := range block.Segments {
		s := block.Segments[i]
		buff.Write(s.Bytes)
	}
	return
}
