package main

import "testing"

// para 是参数
// one 代表第一个参数
type para3 struct {
	s string
}

// ans 是答案
// one 代表第一个答案
type ans3 struct {
	one int
}

var lengthOfLongestSubstringCases = []struct {
	name   string
	input  para3
	except ans3
}{
	{
		"test1",
		para3{"abcabcbb"},
		ans3{3},
	},

	{
		"test2",
		para3{"bbbbb"},
		ans3{1},
	},

	{
		"test3",
		para3{"pwwkew"},
		ans3{3},
	},

	{
		"test4",
		para3{""},
		ans3{0},
	},
}

func TestLengthOfLongestSubstring(t *testing.T) {
	t.Parallel()
	for _, c := range lengthOfLongestSubstringCases {
		t.Run(c.name, func(t *testing.T) {
			actual := lengthOfLongestSubstring(c.input.s)
			if actual != c.except.one {
				t.Errorf("lengthOfLongestSubstring %s test  has fail: input:%v ,except:%v, actual:%v \n", c.name, c.input, c.except.one, actual)
			}
		})
	}
}

func TestLengthOfLongestSubstring2(t *testing.T) {
	t.Parallel()
	for _, c := range lengthOfLongestSubstringCases {
		t.Run(c.name, func(t *testing.T) {
			actual := lengthOfLongestSubstring2(c.input.s)
			if actual != c.except.one {
				t.Errorf("lengthOfLongestSubstring %s test  has fail: input:%v ,except:%v, actual:%v \n", c.name, c.input, c.except.one, actual)
			}
		})
	}
}
