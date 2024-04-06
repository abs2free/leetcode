package main

import "testing"

var romanToIntCases = []struct {
	name   string
	input  string
	except int
}{
	{
		"test1",
		"III",
		3,
	},

	{
		"test2",
		"IV",
		4,
	},

	{
		"test3",
		"IX",
		9,
	},

	{
		"test4",
		"LVIII",
		58,
	},

	{
		"test5",
		"MCMXCIV",
		1994,
	},
	{
		"test6",
		"MMXIV",
		2014,
	},
}

func TestRomanToInt(t *testing.T) {
	t.Parallel()
	for _, c := range romanToIntCases {
		t.Run(c.name, func(t *testing.T) {
			actual := romanToInt(c.input)
			if actual != c.except {
				t.Errorf("romanToInt %s has fail: input:%v ,except:%v, actual:%v \n", c.name, c.input, c.except, actual)
			}
		})
	}
}

func TestRomanToInt2(t *testing.T) {
	t.Parallel()
	for _, c := range romanToIntCases {
		t.Run(c.name, func(t *testing.T) {
			actual := romanToInt2(c.input)
			if actual != c.except {
				t.Errorf("romanToInt %s has fail: input:%v ,except:%v, actual:%v \n", c.name, c.input, c.except, actual)
			}
		})
	}
}
