package main

import "testing"

// para 是参数
// one 代表第一个参数
type para209 struct {
	s   int
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans209 struct {
	one int
}

var minSubArrayLenCases = []struct {
	name   string
	input  para209
	except ans209
}{
	{
		"test1",
		para209{
			7,
			[]int{2, 3, 1, 2, 4, 3},
		},
		ans209{2},
	},

	{
		"test2",
		para209{
			4,
			[]int{1, 4, 4},
		},
		ans209{1},
	},

	{
		"test3",
		para209{
			11,
			[]int{1, 1, 1, 1, 1, 1, 1, 1},
		},
		ans209{0},
	},

	{
		"test4",
		para209{
			11,
			[]int{1, 2, 3, 4, 5},
		},
		ans209{3},
	},

	{
		"test5",
		para209{
			80,
			[]int{10, 5, 13, 4, 8, 4, 5, 11, 14, 9, 16, 10, 20, 8},
		},
		ans209{6},
	},
}

func TestMinSubArrayLen(t *testing.T) {
	t.Parallel()
	for _, c := range minSubArrayLenCases {
		t.Run(c.name, func(t *testing.T) {
			actual := minSubArrayLen(c.input.s, c.input.one)
			if actual != c.except.one {
				t.Errorf("minSubArrayLen %s test  has fail: input:%v ,except:%v, actual:%v \n", c.name, c.input, c.except, actual)
			}
		})
	}
}
