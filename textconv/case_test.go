package textconv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNoCase(t *testing.T) {
	names := [][2]string{
		{
			"camelCase",
			"camel case",
		},
		{
			"CAMELCase",
			"camel case",
		},
	}
	for _, name := range names {
		n := NoCase(name[0], DefaultOptions)
		assert.Equal(t, n, name[1])
	}
}

func TestSplitBytes(t *testing.T) {
	names := []string{
		"camelCase",
		"CAMELCase",
		" CAMELCase ",
	}
	result := [][]string{
		{
			"camel",
			"Case",
		},
		{
			"CAMEL",
			"Case",
		},
		{
			"CAMEL",
			"Case",
		},
	}
	for i, name := range names {
		n := SplitString(name)
		assert.Equal(t, len(n), len(result[i]))
		assert.Equal(t, n, result[i])
	}
}

func TestPascalCase(t *testing.T) {
	names := [][2]string{
		{
			"camelCase",
			"CamelCase",
		},
		{
			"CAMELCase",
			"CamelCase",
		},
		{
			" CAMELCase ",
			"CamelCase",
		},
		{
			"camel 7html",
			"Camel_7html",
		},
	}
	for _, name := range names {
		n := PascalCase(name[0])
		assert.Equal(t, n, name[1])
	}
}

func TestCapitalCase(t *testing.T) {
	names := [][2]string{
		{
			"string",
			"String",
		},
		{
			"dot.case",
			"Dot Case",
		},
		{
			"PascalCase",
			"Pascal Case",
		},
		{
			"version 1.2.10",
			"Version 1 2 10",
		},
	}
	for _, name := range names {
		n := CapitalCase(name[0])
		assert.Equal(t, n, name[1])
	}
}

func TestCamelCase(t *testing.T) {
	names := [][2]string{
		{
			"string",
			"string",
		},
		{
			"dot.case",
			"dotCase",
		},
		{
			"PascalCase",
			"pascalCase",
		},
		{
			"version 1.2.10",
			"version_1_2_10",
		},
	}
	for _, name := range names {
		n := CamelCase(name[0])
		assert.Equal(t, n, name[1])
	}
}

func TestSnakeCase(t *testing.T) {
	names := [][2]string{
		{
			"string",
			"string",
		},
		{
			"dot.case",
			"dot_case",
		},
		{
			"PascalCase",
			"pascal_case",
		},
		{
			"version 1.2.10",
			"version_1_2_10",
		},
	}
	for _, name := range names {
		n := SnakeCase(name[0])
		assert.Equal(t, n, name[1])
	}
}

func TestPathCase(t *testing.T) {
	names := [][2]string{
		{
			"string",
			"string",
		},
		{
			"dot.case",
			"dot/case",
		},
		{
			"PascalCase",
			"pascal/case",
		},
		{
			"version 1.2.10",
			"version/1/2/10",
		},
	}
	for _, name := range names {
		n := PathCase(name[0])
		assert.Equal(t, n, name[1])
	}
}

var ss = " CAMELCase "

func BenchmarkSplitString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SplitString(ss)
	}
}

//var bb = []byte(" CAMELCase ")
//
//func BenchmarkSplitBytes(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		SplitBytes(bb)
//	}
//}
//
//func BenchmarkRegexpSplitBytes(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		RegexpSplitString(ss)
//	}
//}
