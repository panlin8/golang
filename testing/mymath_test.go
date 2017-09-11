package mymath

import (
	"testing"
)

type AddArray struct {
	result int
	num_1  int
	num_2  int
}

func TestAdd(t *testing.T) {
	testdata := [3]AddArray{
		{2, 1, 1},
		{4, 2, 2},
		{6, 3, 3}}

	for _, v := range testdata {
		if v.result != Add(v.num_1, v.num_2) {
			t.Errorf("Add(%d, %d) != %d\n", v.num_1, v.num_2, v.result)
		}
	}
}

func BenchmarkLoops(b *testing.B) {
	var test ForTest
	ptr := &test

	for i := 0; i < b.N; i++ {
		ptr.Loops()
	}
}

func BenchmarkLoopsParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var test ForTest
		ptr := &test

		for pb.Next() {
			ptr.Loops()
		}
	})
}
