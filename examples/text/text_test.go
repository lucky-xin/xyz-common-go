package text

import (
	bytes2 "bytes"
	"github.com/lucky-xin/xyz-common-go/pointer"
	"github.com/lucky-xin/xyz-common-go/text"
	"io"
	"math/rand"
	"os"
	"testing"
	"time"
)

func TestCreateTextFromFile(t *testing.T) {
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

func TestTextToString(t *testing.T) {
	var segments []*text.Segment
	count := 20
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	for range count {
		str := createRandomStr(random.Int() % (128))
		segments = append(segments, &text.Segment{Length: len(str), Bytes: []byte(str)})
	}
	magicByte := pointer.Ptr(byte(8))
	block := &text.Block{MagicByte: magicByte, Segments: segments}
	buff, err := block.ToBuffer()
	if err != nil {
		panic(err)
	}
	block, err = text.FromBuffer(magicByte, &buff)
	println(len(block.Segments))
}

func createRandomStr(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)

	result := make([]byte, length)
	for i := range result {
		result[i] = charset[random.Intn(len(charset))]
	}
	return string(result)
}
