package main

import "testing"

type addBinaryInput struct {
	a string
	b string
}

var addBinaryCases = []struct {
	name   string
	input  addBinaryInput
	except string
}{
	{
		"test1",
		addBinaryInput{"1", "111"},
		"1000",
	},

	{
		"test2",
		addBinaryInput{"110010", "10111"},
		"1001001",
	},

	{

		"test3",
		addBinaryInput{"11", "1"},
		"100",
	},

	{
		"test4",
		addBinaryInput{"1010", "1011"},
		"10101",
	},
}

func TestAddBinary(t *testing.T) {
	t.Parallel()
	for _, c := range addBinaryCases {
		t.Run(c.name, func(t *testing.T) {
			actual := addBinary(c.input.a, c.input.b)
			if actual != c.except {
				t.Errorf("addBinary %s test  has fail: input:%v ,except:%v, actual:%v \n", c.name, c.input, c.except, actual)
			}
		})
	}
}
