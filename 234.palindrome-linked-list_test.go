package main

import "testing"

var palindromeLinkedListCases = []struct {
	name   string
	input  []int
	except bool
}{

	{
		"test1",
		[]int{1, 1, 2, 2, 3, 4, 4, 4},
		false,
	},

	{
		"test2",
		[]int{1, 1, 1, 1, 1, 1},
		true,
	},

	{
		"test3",
		[]int{1, 2, 2, 1, 3},
		false,
	},

	{
		"test4",
		[]int{1},
		true,
	},

	{
		"test5",
		[]int{},
		true,
	},

	{
		"test6",
		[]int{1, 2, 2, 2, 2, 1},
		true,
	},

	{
		"test7",
		[]int{1, 2, 2, 3, 3, 3, 3, 2, 2, 1},
		true,
	},

	{
		"test8",
		[]int{1, 2},
		false,
	},

	{
		"test9",
		[]int{1, 0, 1},
		true,
	},

	{
		"test10",
		[]int{1, 1, 2, 1},
		false,
	},
}

func TestPalindromeLinkedList(t *testing.T) {
	for _, c := range palindromeLinkedListCases {
		t.Run(c.name, func(t *testing.T) {
			actual := isPalindromeLinkedList(Ints2List(c.input))
			if actual != c.except {
				t.Errorf("isPalindromeLinkedList test has fail: input:%v ,except:%v actual:%v \n", c.input, c.except, actual)
			}
		})
	}
}
