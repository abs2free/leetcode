package main

import (
	"testing"
)

func TestMajorityElement(t *testing.T) {
	cases := []struct {
		name    string
		intput []int
		output int
	}{
		{
			"test1",
			[]int{2, 2, 1},
			2,
		},

		{
			"test2",
			[]int{3, 2, 3},
			3,
		},

		{
			"test3",
			[]int{2, 2, 1, 1, 1, 2, 2},
			2,
		},

		{
			"test4",
			[]int{-2147483648},
			-2147483648,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if majorityElement(c.intput) != c.output {
				t.Errorf("majorityElement test has fail: input:%v ,output:%v \n", c.intput, c.output)
			}
		})
	}
}

func BenchmarkMajorityElement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		majorityElement([]int{2, 2, 1, 1, 1, 2, 2})
	}
}
