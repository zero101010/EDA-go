package main

import "testing"

type test struct {
	data   []int
	answer int
}

func TestSomaTabela(t *testing.T) {
	tests := []test{
		test{[]int{1, 2, 3}, 6},
		test{[]int{10, 20, 30}, 60},
		test{[]int{-1, 2, -1}, 0},
	}
	for _, v := range tests {
		z := soma(v.data...)
		if z != v.answer {
			t.Error("Expected:", v.answer, "Got", z)
		}
	}
}
func TestSoma(t *testing.T) {
	test := soma(3, 2, 1)
	result := 6
	if test != result {
		t.Error("Expected:", result, "Got", test)
	}
}

func BenchmarkSoma(b *testing.B) {
	tests := []test{
		test{[]int{1, 2, 3}, 6},
		test{[]int{10, 20, 30}, 60},
		test{[]int{-1, 2, -1}, 0},
	}
	for _, v := range tests {
		soma(v.data...)
	}
}

func TestSomaFalha(t *testing.T) {
	test := soma(3, 2, 1)
	result := 7
	if test == result {
		t.Error("Expected:", result, "Different", test)
	}
}
