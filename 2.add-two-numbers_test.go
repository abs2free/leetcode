package main

import (
	"testing"
)

type addTwoNumbersInput struct {
	l1 []int
	l2 []int
}

var addTwoNumbersCases = []struct {
	name   string
	input  addTwoNumbersInput
	except []int
}{

	{
		"test1",
		addTwoNumbersInput{[]int{}, []int{}},
		[]int{},
	},

	{
		"test2",
		addTwoNumbersInput{[]int{1}, []int{1}},
		[]int{2},
	},

	{
		"test3",
		addTwoNumbersInput{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
		[]int{2, 4, 6, 8},
	},

	{
		"test4",
		addTwoNumbersInput{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		[]int{2, 4, 6, 8, 0, 1},
	},

	{
		"test5",
		addTwoNumbersInput{[]int{1}, []int{9, 9, 9, 9, 9}},
		[]int{0, 0, 0, 0, 0, 1},
	},

	{
		"test6",
		addTwoNumbersInput{[]int{9, 9, 9, 9, 9}, []int{1}},
		[]int{0, 0, 0, 0, 0, 1},
	},

	{
		"test7",
		addTwoNumbersInput{[]int{2, 4, 3}, []int{5, 6, 4}},
		[]int{7, 0, 8},
	},

	{
		"test8",
		addTwoNumbersInput{[]int{1, 8, 3}, []int{7, 1}},
		[]int{8, 9, 3},
	},
}

func TestAddTwoNumbers(t *testing.T) {
	t.Parallel()
	for _, c := range addTwoNumbersCases {
		t.Run(c.name, func(t *testing.T) {
			actual := addTwoNumbers(newByInts(c.input.l1), newByInts(c.input.l2))
			if !equalListNode(newByInts(c.except), actual) {
				t.Errorf("addTwoNumbers %s test  has fail: input:%v ,except:%v, actual:%v \n", c.name, c.input, c.except, formatInts(actual))
			}
		})
	}
}
