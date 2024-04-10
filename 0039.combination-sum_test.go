package main

import (
	"reflect"
	"testing"
)

// para 是参数
// one 代表第一个参数
type para39 struct {
	n []int
	k int
}

// ans 是答案
// one 代表第一个答案
type ans39 struct {
	one [][]int
}

var combinationSumCases = []struct {
	name   string
	input  para39
	except ans39
}{
	{
		"test1",
		para39{[]int{2, 3, 6, 7}, 7},
		ans39{[][]int{{7}, {2, 2, 3}}},
	},

	{
		"test2",
		para39{[]int{2, 3, 5}, 8},
		ans39{[][]int{{3, 5}, {2, 3, 3}, {2, 2, 2, 2}}},
	},
}

func TestCombinationSum(t *testing.T) {
	t.Parallel()
	for _, c := range combinationSumCases {
		t.Run(c.name, func(t *testing.T) {
			actual := combinationSum(c.input.n, c.input.k)
			if !reflect.DeepEqual(actual, c.except.one) {
				t.Errorf("combinationSum %s test  has fail: input:%v ,except:%v, actual:%v \n", c.name, c.input, c.except, actual)
			}
		})
	}
}

func TestCombinationSingle(t *testing.T) {
	c := combinationSumCases[0]
	actual := combinationSum(c.input.n, c.input.k)
	if !reflect.DeepEqual(actual, c.except.one) {
		t.Errorf("combinationSum %s test  has fail: input:%v ,except:%v, actual:%v \n", c.name, c.input, c.except, actual)
	}
}
