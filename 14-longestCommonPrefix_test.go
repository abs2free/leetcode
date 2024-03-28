package main

import "testing"

var longestCommonPrefixCases = []struct {
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
	t.Parallel()
	for _, c := range longestCommonPrefixCases {
		t.Run(c.name, func(t *testing.T) {
			actual := longestCommonPrefix(c.input)
			if actual != c.except {
				t.Errorf("longestCommonPrefix test has fail: input:%v ,except:%v, actual:%v \n", c.input, c.except, actual)
			}
		})
	}
}

func TestLongestCommonPrefix2(t *testing.T) {
	t.Parallel()
	for _, c := range longestCommonPrefixCases {
		t.Run(c.name, func(t *testing.T) {
			actual := longestCommonPrefix2(c.input)
			if actual != c.except {
				t.Errorf("longestCommonPrefix test has fail: input:%v ,except:%v, actual:%v \n", c.input, c.except, actual)
			}
		})
	}
}

func TestLongestCommonPrefix3(t *testing.T) {
	t.Parallel()
	for _, c := range longestCommonPrefixCases {
		t.Run(c.name, func(t *testing.T) {
			actual := longestCommonPrefix3(c.input)
			if actual != c.except {
				t.Errorf("longestCommonPrefix test has fail: input:%v ,except:%v, actual:%v \n", c.input, c.except, actual)
			}
		})
	}
}
