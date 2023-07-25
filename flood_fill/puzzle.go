package main

import "fmt"

type pair struct {
	l int
	r int
}

type Stack[T pair] []T

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

func key(l int, r int) string {
	return fmt.Sprintf("%d-%d", l, r)
}
func inBound(l int, r int, m int, n int) bool {
	return l >= 0 && l <= m && r >= 0 && r <= n
}
func isVisited(visited map[string]bool, sr int, sc int) bool {
	if _, ok := visited[key(sr, sc)]; ok {
		return ok
	}

	return false
}

type Directions [][]int

func initDirections() Directions {
	return [][]int{
		{-1, 0}, // up
		{0, -1}, // left
		{1, 0},  // down
		{0, 1},  // right
	}
}

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	initColor := image[sr][sc]
	s := Stack[pair]{}
	s.Push(pair{l: sr, r: sc})
	visited := map[string]bool{key(sr, sc): true}
	image[sr][sc] = color
	m := len(image) - 1
	n := len(image[0]) - 1

	directions := initDirections()

	for s.IsEmpty() == false {
		curVal, _ := s.Pop()
		sr = curVal.l
		sc = curVal.r

		for _, d := range directions {
			l := sr + d[0]
			r := sc + d[1]
			if inBound(l, r, m, n) && image[l][r] == initColor && isVisited(visited, l, r) == false {
				s.Push(pair{l: l, r: r})
				image[l][r] = color
				visited[key(l, r)] = true
			}
		}
	}

	return image
}
