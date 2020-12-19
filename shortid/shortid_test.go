package shortid

import (
	"github.com/teris-io/shortid"
	"testing"
)

func TestGenerate(t *testing.T) {
	s := Generate()
	t.Log(s)
}

func BenchmarkGenerate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Generate()
	}
}

func BenchmarkGenerate2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = shortid.Generate()
	}
}
