package main

import "testing"

type para35 struct {
	nums   []int
	target int
}

// ans 是答案
// one 代表第一个答案
type ans35 struct {
	one int
}

var searchInsertCases = []struct {
	name   string
	input  para35
	except ans35
}{

	{
		"test1",
		para35{[]int{1, 3, 5, 6}, 5},
		ans35{2},
	},

	{
		"test2",
		para35{[]int{1, 3, 5, 6}, 2},
		ans35{1},
	},

	{
		"test3",
		para35{[]int{1, 3, 5, 6}, 7},
		ans35{4},
	},

	{
		"test4",
		para35{[]int{1, 3, 5, 6}, 0},
		ans35{0},
	},
}

func TestSearchInsert(t *testing.T) {
	t.Parallel()
	for _, c := range searchInsertCases {
		t.Run(c.name, func(t *testing.T) {
			actual := searchInsert(c.input.nums, c.input.target)
			if actual != c.except.one {
				t.Errorf("searchInsert %s test  has fail: input:%v ,except:%v, actual:%v \n", c.name, c.input, c.except, actual)
			}
		})
	}
}

func TestSearchInsert2(t *testing.T) {
	t.Parallel()
	for _, c := range searchInsertCases {
		t.Run(c.name, func(t *testing.T) {
			actual := searchInsert2(c.input.nums, c.input.target)
			if actual != c.except.one {
				t.Errorf("searchInsert %s test  has fail: input:%v ,except:%v, actual:%v \n", c.name, c.input, c.except, actual)
			}
		})
	}
}
