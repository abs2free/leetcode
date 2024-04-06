package main

import "testing"

var isValidParenthesesCases = []struct {
	name   string
	input  string
	except bool
}{
	{
		"test1",
		"()[]{}",
		true,
	},
	{
		"test2",
		"(]",
		false,
	},
	{
		"test3",
		"({[]})",
		true,
	},
	{
		"test4",
		"(){[({[]})]}",
		true,
	},
	{
		"test5",
		"((([[[{{{",
		false,
	},
	{
		"test6",
		"(())]]",
		false,
	},
	{
		"test7",
		"",
		true,
	},
	{
		"test8",
		"[",
		false,
	},
}

func TestIsValidParentheses(t *testing.T) {
	t.Parallel()
	for _, c := range isValidParenthesesCases {
		t.Run(c.name, func(t *testing.T) {
			actual := isValidParentheses(c.input)
			if actual != c.except {
				t.Errorf("isValidParentheses %s test  has fail: input:%v ,except:%v, actual:%v \n", c.name, c.input, c.except, actual)
			}
		})
	}
}
