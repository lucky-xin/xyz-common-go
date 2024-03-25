package main

import (
	"fmt"
	"time"
)

func main() {
	println(time.Now().Format("2006-01-02T15:04:05.999Z07:00"))
	text := "2023-01-29T09:15:09.948+08:00"
	parse, err := time.Parse("2006-01-02T15:04:05.999Z07:00", text)
	if err != nil {
		panic(err)
	}
	println(parse.Hour())
	i := 1000
	println(fmt.Sprint(i))
}
