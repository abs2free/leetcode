package main

import "testing"

// para 是参数
// one 代表第一个参数
type para7 struct {
	one int
}

// ans 是答案
// one 代表第一个答案
type ans7 struct {
	one int
}

var reverseCases = []struct {
	name   string
	input  para7
	except ans7
}{
	{
		"test1",
		para7{321},
		ans7{123},
	},

	{
		"test2",
		para7{-123},
		ans7{-321},
	},

	{
		"test3",
		para7{120},
		ans7{21},
	},

	{
		"test4",
		para7{1534236469},
		ans7{0},
	},
}

func TestReverse(t *testing.T) {
	t.Parallel()
	for _, c := range reverseCases {
		t.Run(c.name, func(t *testing.T) {
			actual := reverse(c.input.one)
			if actual != c.except.one {
				t.Errorf("reverse %s test  has fail: input:%v ,except:%v, actual:%v \n", c.name, c.input, c.except, actual)
			}
		})
	}
}
