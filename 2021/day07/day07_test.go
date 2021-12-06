package day07

import "testing"

func BenchmarkGetInput(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetInput("input")
	}
}

func BenchmarkPart1(b *testing.B) {
	lines := GetInput("input")
	for n := 0; n < b.N; n++ {
		SolvePart1(lines)
	}
}

func BenchmarkPart2(b *testing.B) {
	lines := GetInput("input")
	for n := 0; n < b.N; n++ {
		SolvePart2(lines)
	}
}
func TestSolvePart1(t *testing.T) {
	lines := GetInput("input")
	got := SolvePart1(lines)
	want := 345035

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestSolvePart2(t *testing.T) {
	lines := GetInput("input")
	got := SolvePart2(lines)
	want := 97038163

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
