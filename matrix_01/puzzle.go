package main

import "errors"

type LinkedListNode struct {
	val  state
	next *LinkedListNode
}

type LinkedList struct {
	head *LinkedListNode
	cur  *LinkedListNode
}

func (l *LinkedList) add(s state) {
	node := &LinkedListNode{val: s, next: nil}
	if l.head == nil {
		l.head = node
		l.cur = node
	} else {
		l.cur.next = node
		l.cur = node
	}
}

func (l *LinkedList) remove() (*LinkedListNode, error) {
	if l.head == nil {
		return nil, errors.New("empty")
	}

	el := l.head
	l.head = l.head.next
	if l.head == nil {
		l.cur = nil
	}
	return el, nil
}

func (l *LinkedList) isEmpty() bool {
	return l.head == nil
}

type state struct {
	row   int
	col   int
	steps int
}

func valid(i, j, m, n int) bool {
	return i >= 0 && i < m && j >= 0 && j < n
}

func updateMatrix(mat [][]int) [][]int {
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	m := len(mat)
	n := len(mat[0])

	matrix := mat
	seen := make([][]bool, m)
	for i := 0; i < m; i++ {
		seen[i] = make([]bool, n)
	}
	q := LinkedList{}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			matrix[i][j] = mat[i][j]
			if matrix[i][j] == 0 {
				seen[i][j] = true
				q.add(state{
					row:   i,
					col:   j,
					steps: 0,
				})
			}
		}
	}

	for q.isEmpty() == false {
		s, _ := q.remove()
		for _, d := range directions {
			newRow, newCol, steps := s.val.row+d[0], s.val.col+d[1], s.val.steps

			if valid(newRow, newCol, m, n) {
				isSeen := seen[newRow][newCol]
				if !isSeen {
					seen[newRow][newCol] = true
					q.add(state{
						row:   newRow,
						col:   newCol,
						steps: steps + 1,
					})
					matrix[newRow][newCol] = steps + 1
				}
			}
		}
	}

	return matrix
}
