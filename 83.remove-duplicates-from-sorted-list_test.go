package main

import "testing"

var deleteDuplicatesCases = []struct {
	name   string
	input  []int
	except []int
}{

	{
		"test1",
		[]int{1, 1, 2},
		[]int{1, 2},
	},

	{
		"test2",
		[]int{1, 1, 2, 2, 3, 3, 3},
		[]int{1, 2, 3},
	},

	{
		"test3",
		[]int{1, 1, 1, 1, 1, 1, 1, 1},
		[]int{1},
	},
}

func TestDeleteDuplicates(t *testing.T) {
	t.Parallel()
	for _, c := range deleteDuplicatesCases {
		t.Run(c.name, func(t *testing.T) {
			actual := deleteDuplicates(newByInts(c.input))
			if LinkedListNotEqual(actual, newByInts(c.except)) {
				t.Errorf("deleteDuplicates %s test  has fail: input:%v ,except:%v, actual:%v \n", c.name, c.input, c.except, format2Ints(actual))
			}
		})
	}
}
