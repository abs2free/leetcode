package main

import "testing"

type removeElementsInput struct {
	ll  []int
	val int
}

var removeElementsCases = []struct {
	name   string
	input  removeElementsInput
	except []int
}{
	{
		"test1",
		removeElementsInput{[]int{1, 2, 3, 4, 5}, 1},
		[]int{2, 3, 4, 5},
	},

	{
		"test1",
		removeElementsInput{[]int{1, 2, 3, 4, 5}, 2},
		[]int{1, 3, 4, 5},
	},

	{
		"test1",
		removeElementsInput{[]int{1, 1, 1, 1, 1}, 1},
		[]int{},
	},

	{
		"test1",
		removeElementsInput{[]int{1, 2, 3, 2, 3, 2, 3, 2}, 2},
		[]int{1, 3, 3, 3},
	},

	{
		"test1",
		removeElementsInput{[]int{1, 2, 3, 4, 5}, 5},
		[]int{1, 2, 3, 4},
	},

	{
		"test1",
		removeElementsInput{[]int{}, 5},
		[]int{},
	},

	{
		"test1",
		removeElementsInput{[]int{1, 2, 3, 4, 5}, 10},
		[]int{1, 2, 3, 4, 5},
	},

	{
		"test1",
		removeElementsInput{[]int{1}, 1},
		[]int{},
	},
}

func TestRemoveElements(t *testing.T) {
	t.Parallel()
	for _, c := range removeElementsCases {
		t.Run(c.name, func(t *testing.T) {
			actual := removeElements(Ints2ListNode(c.input.ll), c.input.val)
			if listNodeNotEqual(Ints2ListNode(c.except), actual) {
				t.Errorf("removeElements %s test  has fail: input:%v ,except:%v, actual:%v \n", c.name, c.input, c.except, listNode2Ints(actual))
			}
		})
	}
}

func TestRemoveElements2(t *testing.T) {
	t.Parallel()
	for _, c := range removeElementsCases {
		t.Run(c.name, func(t *testing.T) {
			actual := removeElements2(Ints2ListNode(c.input.ll), c.input.val)
			if listNodeNotEqual(Ints2ListNode(c.except), actual) {
				t.Errorf("removeElements %s test  has fail: input:%v ,except:%v, actual:%v \n", c.name, c.input, c.except, listNode2Ints(actual))
			}
		})
	}
}
