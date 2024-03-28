package main

import "testing"

var longestCommonPrefix_cases = []struct {
	name   string
	input  []string
	except string
}{
	{
		"test1",
		[]string{"flower", "flow", "flight"},
		"fl",
	},
	{
		"test2",
		[]string{"dog", "racecar", "car"},
		"",
	},
	{
		"test3",
		[]string{"ab", "a"},
		"a",
	},
	{
		"test4",
		[]string{"cir", "car"},
		"c",
	},
}

func TestLongestCommonPrefix(t *testing.T) {
	for _, c := range longestCommonPrefix_cases {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			actual := longestCommonPrefix(c.input)
			if actual != c.except {
				t.Errorf("longestCommonPrefix test has fail: input:%v ,except:%v, actual:%v \n", c.input, c.except, actual)
			}
		})
	}
}

func TestLongestCommonPrefix_2(t *testing.T) {
	for _, c := range longestCommonPrefix_cases {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			actual := longestCommonPrefix_2(c.input)
			if actual != c.except {
				t.Errorf("longestCommonPrefix test has fail: input:%v ,except:%v, actual:%v \n", c.input, c.except, actual)
			}
		})
	}
}
