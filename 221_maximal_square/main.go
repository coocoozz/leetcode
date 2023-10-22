package main

import (
	"fmt"
	"math"
)

func main() {
	testCases := []struct {
		matrix   [][]byte
		expected int
	}{
		{
			[][]byte{
				{'1', '0', '1', '0', '0'},
				{'1', '0', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '0', '0', '1', '0'},
			},
			4,
		},
		{
			[][]byte{
				{'0', '1'},
				{'1', '0'},
			},
			1,
		},
		{
			[][]byte{
				{'0'},
			},
			0,
		},
		{
			[][]byte{
				{'0', '0', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
				{'0', '0', '0', '0', '1'},
				{'0', '0', '0', '0', '0'},
			},
			1,
		},
		{
			[][]byte{
				{1},
			},
			1,
		},
	}

	for i, tc := range testCases {
		fmt.Printf("* [Case %d] expected => %d\n", i, tc.expected)
		res := maximalSquare(tc.matrix)
		fmt.Printf("  -- res: %d\n", res)
	}
}

/** Recurrence Relations

  f(i, j) = square bottom right (i, j) 위치에서 가질 수 있는 최대 area
   - matrix(i, j) => i == 0 or j == 0
   - f(i, j)
     - min(f(i-1, j), f(i-1, j-1), f(i, j-1)) + matrix(i, j) => if all of entries are not zero
     - 0 => else..

**/

func maximalSquare(matrix [][]byte) int {
	isAllZero := true
	res := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		res[i] = make([]int, len(matrix[i]))
		for j := 0; j < len(matrix[i]); j++ {
			res[i][j] = int(matrix[i][j]) - 48
			if isAllZero && res[i][j] != 0 {
				isAllZero = false
			}
		}
	}

	maxLengthSquareSide := 0
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			top := res[i-1][j]
			topLeft := res[i-1][j-1]
			left := res[i][j-1]
			cur := res[i][j]
			if top > 0 && topLeft > 0 && left > 0 && cur > 0 {
				res[i][j] = min([]int{top, topLeft, left}) + cur
				if res[i][j] > maxLengthSquareSide {
					maxLengthSquareSide = res[i][j]
				}
			}
		}
	}

	area := maxLengthSquareSide * maxLengthSquareSide
	if area > 0 {
		return area
	} else if !isAllZero {
		return 1
	}
	return 0
}

func max(entries []int) int {
	res := 0
	for _, entry := range entries {
		if res < entry {
			res = entry
		}
	}
	return res
}

func min(entries []int) int {
	res := math.MaxInt
	for _, entry := range entries {
		if res > entry {
			res = entry
		}
	}
	return res
}

func printMatrix(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		row := ""
		for j := 0; j < len(matrix[i]); j++ {
			row += fmt.Sprintf("%3d ", matrix[i][j])
		}
		fmt.Println(row)
	}
}
