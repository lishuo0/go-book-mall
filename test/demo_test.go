package test

import (
	"strings"
	"testing"
)

func BenchmarkStringPlus(b *testing.B) {
	s := ""
	for i := 0; i < b.N; i++ {
		s += "abc"
	}
}

func BenchmarkStringBuilder(b *testing.B) {
	build := strings.Builder{}
	for i := 0; i < b.N; i++ {
		build.WriteString("abc")
	}
}
