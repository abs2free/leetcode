package main

import "testing"

var mySqrtCases = []struct {
	name   string
	input  int
	except int
}{
	{
		"test1",
		4,
		2,
	},

	{
		"test2",
		8,
		2,
	},
	{
		"test3",
		1,
		1,
	},
}

func TestMySqrt(t *testing.T) {
	t.Parallel()
	for _, c := range mySqrtCases {
		t.Run(c.name, func(t *testing.T) {
			actual := mySqrt(c.input)
			if actual != c.except {
				t.Errorf("mySqrt %s test  has fail: input:%v ,except:%v, actual:%v \n", c.name, c.input, c.except, actual)
			}
		})
	}
}
