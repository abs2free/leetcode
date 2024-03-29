package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	st := newStack[int]()
	if !st.IsEmpty() {
		t.Error("this is not a empty stack")
	}

	st.Push(1)
	v := st.Pop()
	if v != 1 {
		t.Error("this is not a empty stack")
	}

	if !st.IsEmpty() {
		t.Errorf("this is not a empty stack")
	}
}
