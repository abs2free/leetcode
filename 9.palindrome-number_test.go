package main

import (
	"testing"
)

var isPalindromeCases = []struct {
	name   string
	input  int
	except bool
}{

	{
		"test1",
		121,
		true,
	},

	{
		"test2",
		-121,
		false,
	},

	{
		"test3",
		10,
		false,
	},

	{
		"test4",
		321,
		false,
	},

	{
		"test5",
		-123,
		false,
	},

	{
		"test6",
		120,
		false,
	},

	{
		"test7",
		153423469,
		false,
	},
}

func TestIsPalindromet(t *testing.T) {
	t.Parallel()
	for _, c := range isPalindromeCases {
		t.Run(c.name, func(t *testing.T) {
			actual := isPalindrome(c.input)
			if actual != c.except {
				t.Errorf("isPalindrome test has fail: input:%v ,except:%v actual:%v \n", c.input, c.except, actual)
			}
		})
	}
}

func TestIsPalindrometString(t *testing.T) {
	t.Parallel()
	for _, c := range isPalindromeCases {
		t.Run(c.name, func(t *testing.T) {
			actual := isPalindromeString(c.input)
			if actual != c.except {
				t.Errorf("isPalindromeString  test has fail: input:%v ,except:%v actual:%v \n", c.input, c.except, actual)
			}
		})
	}
}
