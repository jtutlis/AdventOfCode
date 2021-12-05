package day05

import "testing"

func BenchmarkGetLines(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetLines("input")
	}
}

func BenchmarkSolveFast(b *testing.B) {
	lines := GetLines("input")
	for n := 0; n < b.N; n++ {
		SolveFast(lines, true, false)
	}
}
