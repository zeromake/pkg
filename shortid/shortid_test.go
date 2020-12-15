package shortid

import "testing"

func TestGenerate(t *testing.T) {
	s := Generate()
	t.Log(s)
}

func BenchmarkGenerate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Generate()
	}
}
