package main

func listNodeNotEqual(l1, l2 *ListNode) bool {
	if l1 == nil && l2 == nil {
		return false
	}

	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return true
		}
		l1 = l1.Next
		l2 = l2.Next
	}

	if l1 != nil || l2 != nil {
		return true
	}

	return false
}

func sliceClone[T any](in []T) []T {
	res := make([]T, len(in))
	copy(res, in)
	return res
}
