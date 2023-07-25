package main

import (
	"testing"
)

func TestMatrix1(t *testing.T) {
	updateMatrix([][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	})
}
func TestMatrix2(t *testing.T) {
	updateMatrix([][]int{
		{0, 0, 0},
		{0, 1, 0},
		{1, 1, 1},
	})
}
func TestMatrix3(t *testing.T) {
	updateMatrix([][]int{
		{0}, {0}, {0}, {0}, {0},
	})
}
