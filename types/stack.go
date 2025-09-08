package types

// Stack — стек для любого типа T
type Stack[T any] struct {
	elements []T
}

// NewStack создаёт новый стек
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{elements: []T{}}
}

// Push — добавляет элемент в верх стека
func (s *Stack[T]) Push(value T) {
	s.elements = append(s.elements, value)
}

// Pop — извлекает верхний элемент стека
// возвращает элемент и true, если стек не пуст
func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if len(s.elements) == 0 {
		return zero, false
	}
	lastIndex := len(s.elements) - 1
	val := s.elements[lastIndex]
	s.elements = s.elements[:lastIndex]
	return val, true
}

// Peek — возвращает верхний элемент без удаления
func (s *Stack[T]) Peek() (T, bool) {
	var zero T
	if len(s.elements) == 0 {
		return zero, false
	}
	return s.elements[len(s.elements)-1], true
}

// IsEmpty — проверяет, пустой ли стек
func (s *Stack[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

// Size — возвращает количество элементов
func (s *Stack[T]) Size() int {
	return len(s.elements)
}
