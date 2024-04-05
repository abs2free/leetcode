package main

import "testing"

// para 是参数
// one 代表第一个参数
type para8 struct {
	one string
}

// ans 是答案
// one 代表第一个答案
type ans8 struct {
	one int
}

var myAtoiCases = []struct {
	name   string
	input  para8
	except ans8
}{

	{
		"test1",
		para8{"42"},
		ans8{42},
	},

	{
		"test2",
		para8{"   -42"},
		ans8{-42},
	},

	{
		"test3",
		para8{"4193 with words"},
		ans8{4193},
	},

	{
		"test4",
		para8{"words and 987"},
		ans8{0},
	},

	{
		"test5",
		para8{"-91283472332"},
		ans8{-2147483648},
	},

	{
		"test6",
		para8{"+1"},
		ans8{1},
	},

	{
		"test7",
		para8{"+-12"},
		ans8{0},
	},
	{
		"test8",
		para8{""},
		ans8{0},
	},

	{
		"test9",
		para8{"9223372036854775808"},
		ans8{2147483647},
	},
}

func TestMyAtoi(t *testing.T) {
	t.Parallel()
	for _, c := range myAtoiCases {
		t.Run(c.name, func(t *testing.T) {
			actual := myAtoi(c.input.one)
			if actual != c.except.one {
				t.Errorf("myAtoi %s test  has fail: input:%v ,except:%v, actual:%v \n", c.name, c.input, c.except, actual)
			}
		})
	}
}
