package main

import "testing"

// para 是参数
// one 代表第一个参数
type para11 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans11 struct {
	one int
}

var maxAreaCases = []struct {
	name   string
	input  para11
	except ans11
}{
	{
		"test1",
		para11{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}},
		ans11{49},
	},

	{
		"test2",
		para11{[]int{1, 1}},
		ans11{1},
	},
}

func TestMaxArea(t *testing.T) {
	for _, c := range maxAreaCases {
		t.Run(c.name, func(t *testing.T) {
			actual := maxArea(c.input.one)
			if actual != c.except.one {
				t.Errorf("maxArea %s test  has fail: input:%v ,except:%v, actual:%v \n", c.name, c.input, c.except, actual)
			}
		})
	}
}
