package main

import "testing"

var removeDuplicatesCases = []struct {
	name   string
	input  []int
	except int
}{
	{
		"test1",
		[]int{1, 1, 2},
		2,
	},

	{
		"test2",
		[]int{0, 0, 1, 1, 1, 1, 2, 3, 4, 4},
		5,
	},

	{
		"test3",
		[]int{0, 0, 0, 0, 0},
		1,
	},

	{
		"test5",
		[]int{1},
		1,
	},
}

func TestRemoveDuplicates(t *testing.T) {
	t.Parallel()
	for _, c := range removeDuplicatesCases {
		t.Run(c.name, func(t *testing.T) {
			in := sliceClone[int](c.input)
			actual := removeDuplicates(in)
			if actual != c.except {
				t.Errorf("removeDuplicates %s test  has fail: input:%v ,except:%v, actual:%v \n", c.name, c.input, c.except, actual)
			}
		})
	}
}

func TestRemoveDuplicates2(t *testing.T) {
	t.Parallel()
	for _, c := range removeDuplicatesCases {
		t.Run(c.name, func(t *testing.T) {
			in := sliceClone[int](c.input)
			actual := removeDuplicates2(in)
			if actual != c.except {
				t.Errorf("removeDuplicates %s test  has fail: input:%v ,except:%v, actual:%v \n", c.name, c.input, c.except, actual)
			}
		})
	}
}
