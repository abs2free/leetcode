package stack

type stack[T any] struct {
	Values []T
}

func newStack[T any]() *stack[T] {
	return &stack[T]{}
}

func (s *stack[T]) IsEmpty() bool {
	return len(s.Values) == 0
}

func (s *stack[T]) Push(v T) {
	s.Values = append(s.Values, v)
}

func (s *stack[T]) Pop() T {
	n := len(s.Values)
	v := s.Values[n-1]
	s.Values = s.Values[:n-1]
	return v
}
