package data

import (
	"math/rand"
	"time"
)

var (
	Data       []string
	dataLength int
)

func Random() string       { return Data[getRandom(dataLength)] }
func Get(index int) string { return Data[index%dataLength] }
func getRandom(n int) int  { return rand.New(rand.NewSource(time.Now().Unix())).Intn(n) }
