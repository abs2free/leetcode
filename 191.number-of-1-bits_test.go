package main

import "testing"

var hammingWeightCases = []struct {
	name   string
	input  int
	except int
}{

	{
		"test1",
		11,
		3,
	},

	{
		"test2",
		128,
		1,
	},

	{
		"test3",
		2147483645,
		30,
	},
}

func TestHammingWeight(t *testing.T) {
	t.Parallel()
	for _, c := range hammingWeightCases {
		t.Run(c.name, func(t *testing.T) {
			actual := hammingWeight(c.input)
			if actual != c.except {
				t.Errorf("hammingWeight %s test  has fail: input:%v ,except:%v, actual:%v \n", c.name, c.input, c.except, actual)
			}
		})
	}
}

func TestHammingWeight2(t *testing.T) {
	t.Parallel()
	for _, c := range hammingWeightCases {
		t.Run(c.name, func(t *testing.T) {
			actual := hammingWeight2(c.input)
			if actual != c.except {
				t.Errorf("hammingWeight %s test  has fail: input:%v ,except:%v, actual:%v \n", c.name, c.input, c.except, actual)
			}
		})
	}
}
