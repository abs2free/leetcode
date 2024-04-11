package main

import (
	"reflect"
	"testing"
)

// para 是参数
// one 代表第一个参数
type para78 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans78 struct {
	one [][]int
}

var subsetsCases = []struct {
	name   string
	input  para78
	except ans78
}{
	{
		"test1",
		para78{[]int{}},
		ans78{[][]int{nil}},
	},

	//	{
	//		"test2",
	//		para78{[]int{1, 2, 3}},
	//		ans78{[][]int{{}, {1}, {2}, {3}, {1, 2}, {2, 3}, {1, 3}, {1, 2, 3}}},
	//	},
}

func TestSubsets(t *testing.T) {
	for _, c := range subsetsCases {
		t.Run(c.name, func(t *testing.T) {
			actual := subsets(c.input.one)
			if !reflect.DeepEqual(c.except.one, actual) {
				t.Errorf("subsets %s test  has fail: input:%v ,except:%v, actual:%v \n", c.name, c.input, c.except.one, actual)
			}
		})
	}
}

func TestSubsetsSingle(t *testing.T) {
	c := subsetsCases[0]
	actual := subsets(c.input.one)
	if !reflect.DeepEqual(c.except.one, actual) {
		t.Errorf("subsets %s test  has fail: input:%v ,except:%v, actual:%v \n", c.name, c.input, c.except.one, actual)
	}
}
