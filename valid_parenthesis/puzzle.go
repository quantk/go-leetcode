package main

import "fmt"

type Stack[T string | int] []T

func (s *Stack[T]) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack[T]) Push(val T) {
	*s = append(*s, val)
}

func (s *Stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, fmt.Errorf("empty")
	} else {
		index := len(*s) - 1
		value := (*s)[index]
		*s = (*s)[:index]
		return value, nil
	}
}

func isOpened(val rune) bool {
	switch val {
	case '(', '{', '[':
		return true
	default:
		return false
	}
}

func isValid(s string) bool {
	if s == "" {
		return true
	}

	if len(s)%2 != 0 {
		return false
	}

	stack := Stack[string]{}

	brMap := map[rune]string{
		')': "(",
		'}': "{",
		']': "[",
	}

	for _, val := range s {
		if isOpened(val) {
			stack.Push(string(val))
		} else {
			lastVal, _ := stack.Pop()
			if lastVal != brMap[val] {
				return false
			}
		}
	}

	return stack.IsEmpty()
}
