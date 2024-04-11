package main

import (
	"slices"
	"testing"
)

var plusOneCases = []struct {
	name   string
	input  []int
	except []int
}{

	{
		"test1",
		[]int{1, 2, 3},
		[]int{1, 2, 4},
	},

	{
		"test4",
		[]int{4, 3, 2, 1},
		[]int{4, 3, 2, 2},
	},

	{
		"test2",
		[]int{9, 9},
		[]int{1, 0, 0},
	},

	{
		"test3",
		[]int{0},
		[]int{1},
	},
}

func TestPlusOne(t *testing.T) {
	t.Parallel()
	for _, c := range plusOneCases {
		t.Run(c.name, func(t *testing.T) {
			actual := plusOne(c.input)
			if !slices.Equal(actual, c.except) {
				t.Errorf("plusOne %s test  has fail: input:%v ,except:%v, actual:%v \n", c.name, c.input, c.except, actual)
			}
		})
	}
}
