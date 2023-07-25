package main

import "testing"

func TestQueue(t *testing.T) {
	q := Constructor()
	q.Push(1)
	q.Push(2)
	if q.Peek() != 1 {
		t.Errorf("failed peek")
	}

	if q.Pop() != 1 {
		t.Errorf("failed pop")
	}

	if q.Empty() == true {
		t.Errorf("failed empty")
	}
}
