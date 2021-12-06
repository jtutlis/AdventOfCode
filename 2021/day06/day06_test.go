package day06

import "testing"

func BenchmarkGetInput(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetInput()
	}
}

func BenchmarkPart1(b *testing.B) {
	lines := GetInput()
	for n := 0; n < b.N; n++ {
		Solve(lines, 80)
	}
}

func BenchmarkPart2(b *testing.B) {
	lines := GetInput()
	for n := 0; n < b.N; n++ {
		Solve(lines, 256)
	}
}

func BenchmarkParseAndSolve(b *testing.B) {
	for n := 0; n < b.N; n++ {
		lines := GetInput()
		Solve(lines, 256)
	}
}

func TestSolvePart1(t *testing.T) {
	lines := GetInput()
	got := Solve(lines, 80)
	want := 371379

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestSolvePart2(t *testing.T) {
	lines := GetInput()
	got := Solve(lines, 256)
	want := 1674303997472

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
