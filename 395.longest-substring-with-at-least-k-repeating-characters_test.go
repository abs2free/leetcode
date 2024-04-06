package main

import "testing"

// para 是参数
// one 代表第一个参数
type para395 struct {
	s string
	k int
}

// ans 是答案
// one 代表第一个答案
type ans395 struct {
	one int
}

var longestSubstringCases = []struct {
	name   string
	input  para395
	except ans395
}{
	{
		"test1",
		para395{"aaabb", 3},
		ans395{3},
	},

	{
		"test2",
		para395{"ababbc", 2},
		ans395{5},
	},
}

func TestLongestSubstring(t *testing.T) {
	t.Parallel()
	for _, c := range longestSubstringCases {
		t.Run(c.name, func(t *testing.T) {
			actual := longestSubstring(c.input.s, c.input.k)
			if actual != c.except.one {
				t.Errorf("longestSubstring %s test  has fail: input:%v ,except:%v, actual:%v \n", c.name, c.input, c.except, actual)
			}
		})
	}
}
