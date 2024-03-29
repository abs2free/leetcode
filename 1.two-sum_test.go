package main

import (
	"slices"
	"testing"
)

var twoSumCases = []struct {
	name  string
	input struct {
		nums   []int
		target int
	}
	except []int
}{
	{
		"test1",
		struct {
			nums   []int
			target int
		}{[]int{3, 2, 4}, 6},
		[]int{1, 2},
	},

	{
		"test2",
		struct {
			nums   []int
			target int
		}{[]int{3, 2, 4}, 5},
		[]int{0, 1},
	},

	{
		"test3",
		struct {
			nums   []int
			target int
		}{[]int{0, 8, 7, 3, 3, 4, 2}, 11},
		[]int{1, 3},
	},

	{
		"test4",
		struct {
			nums   []int
			target int
		}{[]int{0, 1}, 1},
		[]int{0, 1},
	},

	{
		"test5",
		struct {
			nums   []int
			target int
		}{[]int{0, 3}, 5},
		[]int{},
	},
}

func TestTwoSum(t *testing.T) {
	t.Parallel()
	for _, c := range twoSumCases {
		t.Run(c.name, func(t *testing.T) {
			actual := twoSum(c.input.nums, c.input.target)
			if !slices.Equal(actual, c.except) {
				t.Errorf("twoSum %s has fail: input:%v ,except:%v, actual:%v \n", c.name, c.input, c.except, actual)
			}
		})
	}
}
