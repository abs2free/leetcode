package main

import "testing"

// para 是参数
// one 代表第一个参数
type para219 struct {
	one []int
	k   int
}

// ans 是答案
// one 代表第一个答案
type ans219 struct {
	one bool
}

var containsNearbyDuplicateCases = []struct {
	name   string
	input  para219
	except ans219
}{

	{
		"test1",
		para219{[]int{1, 2, 3, 1}, 3},
		ans219{true},
	},

	{
		"test2",
		para219{[]int{1, 0, 0, 1}, 1},
		ans219{true},
	},

	{
		"test3",
		para219{[]int{1, 2, 3, 1, 2, 3}, 2},
		ans219{false},
	},
	{
		"test4",
		para219{[]int{1, 2}, 2},
		ans219{false},
	},
	{
		"test5",
		para219{[]int{1, 2, 1}, 0},
		ans219{false},
	},

	{
		"test6",
		para219{[]int{0, 1, 2, 3, 2, 5}, 3},
		ans219{true},
	},
	{
		"test7",
		para219{[]int{1, 2, 2, 3}, 3},
		ans219{true},
	},
}

func TestContainsNearbyDuplicate(t *testing.T) {
	t.Parallel()
	for _, c := range containsNearbyDuplicateCases {
		t.Run(c.name, func(t *testing.T) {
			actual := containsNearbyDuplicate(c.input.one, c.input.k)
			if actual != c.except.one {
				t.Errorf("containsNearbyDuplicate %s test  has fail: input:%v ,except:%v, actual:%v \n", c.name, c.input, c.except, actual)
			}
		})
	}
}
