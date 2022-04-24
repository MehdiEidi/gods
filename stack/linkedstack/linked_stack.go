package linkedstack

import "github.com/MehdiEidi/gods/linkedlist/singly"

type LinkedStack[T any] struct {
	data singly.LinkedList[T]
}

// New constructs and returns an empty linked stack.
func New[T any]() *LinkedStack[T] {
	return &LinkedStack[T]{}
}

// Push adds an element on top of the stack.
func (s *LinkedStack[T]) Push(data T) {
	s.data.AddFirst(data)
}

// Pop removes the top element and returns it. It returns false if the stack was empty.
func (s *LinkedStack[T]) Pop() (T, bool) {
	return s.data.RemoveFirst()
}

// Top returns the top element of the stack. It returns false if the stack was empty.
func (s *LinkedStack[T]) Top() (T, bool) {
	return s.data.First()
}

// Size returns the size of the stack.
func (s *LinkedStack[T]) Size() int {
	return s.data.Size
}

// IsEmpty returns true if the stack doesn't have any elements.
func (s *LinkedStack[T]) IsEmpty() bool {
	return s.data.IsEmpty()
}

// String returns the string representation of the stack.
func (s *LinkedStack[T]) String() string {
	return s.data.String()
}
