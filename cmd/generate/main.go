package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"text/template"
)

func main() {
	bts, err := ioutil.ReadFile("data/data.txt")
	if err != nil {
		panic(err)
	}

	arr := bytes.Split(bytes.ReplaceAll(bts, []byte("\r"), []byte("\n")), []byte("\n"))
	data := Data{Text: make([]string, 0, len(arr)), Length: 0}

	for _, v := range arr {
		if len(v) < 1 {
			continue
		}
		data.Text = append(data.Text, string(v))
	}
	data.Length = len(data.Text)

	var file *os.File
	if file, err = os.OpenFile("data/data.gen.go", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666); err != nil {
		panic(err)
	}

	if err = template.Must(template.New("queue").Parse(tpl)).Execute(file, data); err != nil {
		panic(err)
	}
}

type Data struct {
	Text   []string
	Length int
}

const (
	tpl = `// Code generated, DO NOT EDIT.

package data

var (
	Data = []string{
{{range .Text}}"{{.}}",{{end}}
	}
	dataLength = {{.Length}}
)
`
)
