package print_2D_clockwise

import "testing"

func TestPrintSquare(t *testing.T) {
	demo1 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}               // 3*3
	demo2 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10, 11, 12}} // 3*4
	demo3 := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}   // 4*3
	PrintSquare(demo1)
	PrintSquare(demo2)
	PrintSquare(demo3)
}
