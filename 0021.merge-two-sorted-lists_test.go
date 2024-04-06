package main

import "testing"

type mergeTwoListsInput struct {
	list1 []int
	list2 []int
}

var mergeTwoListsCases = []struct {
	name   string
	input  mergeTwoListsInput
	except []int
}{
	{
		"test1",
		mergeTwoListsInput{[]int{}, []int{}},
		[]int{},
	},

	{
		"test2",
		mergeTwoListsInput{[]int{1}, []int{1}},
		[]int{1, 1},
	},

	{
		"test3",
		mergeTwoListsInput{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
		[]int{1, 1, 2, 2, 3, 3, 4, 4},
	},

	{
		"test4",
		mergeTwoListsInput{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		[]int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5},
	},

	{
		"test5",
		mergeTwoListsInput{[]int{1}, []int{9, 9, 9, 9, 9}},
		[]int{1, 9, 9, 9, 9, 9},
	},

	{
		"test6",
		mergeTwoListsInput{[]int{9, 9, 9, 9, 9}, []int{1}},
		[]int{1, 9, 9, 9, 9, 9},
	},

	{
		"test7",
		mergeTwoListsInput{[]int{2, 3, 4}, []int{4, 5, 6}},
		[]int{2, 3, 4, 4, 5, 6},
	},

	{
		"test8",
		mergeTwoListsInput{[]int{1, 3, 8}, []int{1, 7}},
		[]int{1, 1, 3, 7, 8},
	},
}

func TestMergeTwoLists(t *testing.T) {
	t.Parallel()
	for _, c := range mergeTwoListsCases {
		t.Run(c.name, func(t *testing.T) {
			actual := mergeTwoLists(Ints2ListNode(c.input.list1), Ints2ListNode(c.input.list2))
			if listNodeNotEqual(actual, Ints2ListNode(c.except)) {
				t.Errorf("mergeTwoLists %s test  has fail: input:%v ,except:%v, actual:%v \n", c.name, c.input, c.except, listNode2Ints(actual))
			}
		})
	}
}
