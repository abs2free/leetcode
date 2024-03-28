package main

import (
	"testing"
)

var majorityElementCases = []struct {
	name   string
	input  []int
	except int
}{
	{
		"test1",
		[]int{2, 2, 1},
		2,
	},

	{
		"test2",
		[]int{3, 2, 3},
		3,
	},

	{
		"test3",
		[]int{2, 2, 1, 1, 1, 2, 2},
		2,
	},

	{
		"test4",
		[]int{-2147483648},
		-2147483648,
	},
}

func TestMajorityElement(t *testing.T) {
	t.Parallel()
	for _, c := range majorityElementCases {
		t.Run(c.name, func(t *testing.T) {
			actual := majorityElement(c.input)
			if actual != c.except {
				t.Errorf("majorityElement test has fail: input:%v ,except:%v actual:%v \n", c.input, c.except, actual)
			}
		})
	}
}

func TestMajorityElement2(t *testing.T) {
	t.Parallel()
	for _, c := range majorityElementCases {
		t.Run(c.name, func(t *testing.T) {
			actual := majorityElement2(c.input)
			if actual != c.except {
				t.Errorf("majorityElement test has fail: input:%v ,except:%v actual:%v \n", c.input, c.except, actual)
			}
		})
	}
}
