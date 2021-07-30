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

	arr := bytes.Split(bts, []byte("\n"))
	data := Data{Text: make([]string, 0, len(arr)), Length: 0}

	for _, v := range arr {
		if len(v) < 1 {
			continue
		}
		data.Text = append(data.Text, string(v))
		data.Length++
	}

	var file *os.File
	if file, err = os.OpenFile("data/data.go", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666); err != nil {
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

import (
	"math/rand"
	"time"
)

func Random() string       { return Data[getRandom(dataLength)] }
func Get(index int) string { return Data[index%dataLength] }
func getRandom(n int) int  { return rand.New(rand.NewSource(time.Now().Unix())).Intn(n) }

var (
	Data = []string{
{{range .Text}}"{{.}}",{{end}}
	}
	dataLength = {{.Length}}
)
`
)
