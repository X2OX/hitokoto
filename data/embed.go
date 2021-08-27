package data

import (
	"bytes"
	"embed"
)

//go:embed data.txt
var content embed.FS

func init() {
	data, err := content.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}

	arr := bytes.Split(bytes.ReplaceAll(data, []byte("\r"), []byte("\n")), []byte("\n"))
	Data = make([]string, 0, len(arr))

	for _, v := range arr {
		if len(v) < 1 {
			continue
		}
		Data = append(Data, string(v))
	}
	dataLength = len(Data)
}
