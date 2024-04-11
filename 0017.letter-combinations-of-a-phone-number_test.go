package main

import (
	"slices"
	"testing"
)

// para 是参数
// one 代表第一个参数
type para17 struct {
	s string
}

// ans 是答案
// one 代表第一个答案
type ans17 struct {
	one []string
}

var letterCombinationsCases = []struct {
	name   string
	input  para17
	except ans17
}{
	{
		"test1",
		para17{"23"},
		ans17{[]string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
	},
}

func TestLetterCombinations(t *testing.T) {
	for _, c := range letterCombinationsCases {
		t.Run(c.name, func(t *testing.T) {
			actual := letterCombinations(c.input.s)
			if !slices.Equal(actual, c.except.one) {
				t.Errorf("letterCombinations %s test  has fail: input:%v ,except:%v, actual:%v \n", c.name, c.input, c.except, actual)
			}
		})
	}
}

func TestLetterCombinationsSingle(t *testing.T) {
	c := letterCombinationsCases[0]
	actual := letterCombinations(c.input.s)
	if !slices.Equal(actual, c.except.one) {
		t.Errorf("letterCombinations %s test  has fail: input:%v ,except:%v, actual:%v \n", c.name, c.input, c.except, actual)
	}
}
