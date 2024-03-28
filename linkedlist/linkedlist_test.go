package linkedlist

import (
	"reflect"
	"testing"
)

func list2Ints(head *ListNode) []int {
	var nums []int
	curr := head
	for curr != nil {
		nums = append(nums, curr.Val)
		curr = curr.Next
	}
	return nums
}

func TestInts2List(t *testing.T) {
	cases := []struct {
		name   string
		input  []int
		except []int
	}{
		{
			"test1",
			[]int{1, 2, 3, 4, 5, 4},
			[]int{1, 2, 3, 4, 5, 4},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ls := ints2List(c.input)
			actual := list2Ints(ls)
			if !reflect.DeepEqual(actual, c.except) {
				t.Errorf("TestInts2List has fail: input:%v ,except:%v actual:%v \n", c.input, c.except, actual)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	cases := []struct {
		name   string
		input  []int
		except []int
	}{
		{
			"test1",
			[]int{1, 2, 3, 4, 5, 4},
			[]int{1, 2, 3, 4, 5, 4},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ls := reverse(ints2List(c.input))
			actual := list2Ints(ls)
			if !reflect.DeepEqual(actual, c.except) {
				t.Errorf("TestReverse has fail: input:%v ,except:%v actual:%v \n", c.input, c.except, actual)
			}
		})
	}
}

func TestClone(t *testing.T) {
	cases := []struct {
		name   string
		input  []int
		except []int
	}{
		{
			"test1",
			[]int{1, 2, 3, 4, 5, 4},
			[]int{1, 2, 3, 4, 5, 4},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ls := ints2List(c.input)
			actual := clone(ls)
			if !reflect.DeepEqual(actual, c.except) {
				t.Errorf("TestClone has fail: input:%v ,except:%v actual:%v \n", c.input, c.except, actual)
			}
		})
	}
}
