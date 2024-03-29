package text

import (
	bytes2 "bytes"
	"github.com/lucky-xin/xyz-common-go/text"
	"io"
	"os"
	"testing"
)

func TestText(t *testing.T) {
	f, err := os.Open("eascs_dev.lic")
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}(f)
	if err != nil {
		return
	}
	payload, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	buffer := bytes2.NewBuffer(payload)
	//var magicByte byte = 0x0
	block, err := text.FromBuffer(nil, buffer)
	if err != nil {
		panic(err)
	}
	for i := range block.Segments {
		segment := block.Segments[i]
		println(segment.Length, "---", string(segment.Bytes))
	}
}
