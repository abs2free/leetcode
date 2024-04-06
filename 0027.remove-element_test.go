package main

import "testing"

// para 是参数
// one 代表第一个参数
type para27 struct {
	one []int
	two int
}

// ans 是答案
// one 代表第一个答案
type ans27 struct {
	one int
}

var removeElementCases = []struct {
	name   string
	input  para27
	except ans27
}{
	{
		"test1",
		para27{[]int{1, 0, 1}, 1},
		ans27{1},
	},

	{
		"test2",
		para27{[]int{0, 1, 0, 3, 0, 12}, 0},
		ans27{3},
	},

	{
		"test3",
		para27{[]int{0, 1, 0, 3, 0, 0, 0, 0, 1, 12}, 0},
		ans27{4},
	},

	{
		"test4",
		para27{[]int{0, 0, 0, 0, 0}, 0},
		ans27{0},
	},

	{
		"test5",
		para27{[]int{1}, 1},
		ans27{0},
	},

	{
		"test6",
		para27{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2},
		ans27{5},
	},
	{
		"test7",
		para27{[]int{2}, 3},
		ans27{1},
	},
}

func TestRemoveElement(t *testing.T) {
	for _, c := range removeElementCases {
		t.Run(c.name, func(t *testing.T) {
			actual := removeElement(c.input.one, c.input.two)
			if actual != c.except.one {
				t.Errorf("removeElement %s test  has fail: input:%v ,except:%v, actual:%v \n", c.name, c.input, c.except, actual)
			}
		})
	}
}

func TestRemoveElement2(t *testing.T) {
	for _, c := range removeElementCases {
		t.Run(c.name, func(t *testing.T) {
			actual := removeElement2(c.input.one, c.input.two)
			if actual != c.except.one {
				t.Errorf("removeElement %s test  has fail: input:%v ,except:%v, actual:%v \n", c.name, c.input, c.except, actual)
			}
		})
	}
}
func TestRemoveElement3(t *testing.T) {
	for _, c := range removeElementCases {
		t.Run(c.name, func(t *testing.T) {
			actual := removeElement3(c.input.one, c.input.two)
			if actual != c.except.one {
				t.Errorf("removeElement %s test  has fail: input:%v ,except:%v, actual:%v \n", c.name, c.input, c.except, actual)
			}
		})
	}
}
