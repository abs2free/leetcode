package main

import "testing"

var lengthOfLastWordCases = []struct {
	name   string
	input  string
	except int
}{
	{
		"test1",
		"Hello World",
		5,
	},

	{
		"test2",
		"   fly me   to   the moon  ",
		4,
	},

	{
		"test3",
		"luffy is still joyboy",
		6,
	},
}

func TestLengthOfLastWord(t *testing.T) {
	t.Parallel()
	for _, c := range lengthOfLastWordCases {
		t.Run(c.name, func(t *testing.T) {
			actual := lengthOfLastWord(c.input)
			if actual != c.except {
				t.Errorf("lengthOfLastWord %s test  has fail: input:%v ,except:%v, actual:%v \n", c.name, c.input, c.except, actual)
			}
		})
	}
}
