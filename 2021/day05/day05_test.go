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

func TestSolveFastPart2(t *testing.T) {
	lines := GetLines("input")
	got := SolveFast(lines, true, false)
	want := 22116

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestSolveFastPart1(t *testing.T) {
	lines := GetLines("input")
	got := SolveFast(lines, false, false)
	want := 6225

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
