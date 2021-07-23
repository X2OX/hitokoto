package data

import (
	"testing"
)

func TestFunc(t *testing.T) {
	switch {
	case len(Random()) == 0, len(Get(0)) == 0:
		t.Fail()
	}
	for _, v := range Data {
		if len(v) == 0 {
			t.Fail()
		}
	}
}

func BenchmarkRandom(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Random()
	}
}
func BenchmarkGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Get(i)
	}
}
