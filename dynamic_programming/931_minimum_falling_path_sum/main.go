package main

import (
	"fmt"
	"leetcode/utils"
	"math"
)

func main() {
	testCases := []struct {
		matrix   [][]int
		expected int
	}{
		{
			[][]int{
				{2, 1, 3},
				{6, 5, 4},
				{7, 8, 9},
			},
			13,
		},
		{
			[][]int{
				{-19, 57},
				{-40, -5},
			},
			-59,
		},
		{
			[][]int{
				{-48},
			},
			-48,
		},
	}

	for i, tc := range testCases {
		fmt.Printf("[Case%d] expected: %d\n", i, tc.expected)
		var res int

		elapsed := utils.ElapsedTime(func() {
			res = minFallingPathSum(tc.matrix)
		})
		fmt.Printf("  - res: %d, elapsed: %v\n", res, elapsed)
	}

}

func minFallingPathSum(matrix [][]int) int {
	if len(matrix) == 1 && len(matrix[0]) == 1 {
		return matrix[0][0]
	}

	res := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		res[i] = make([]int, len(matrix[i]))
	}

	ret := math.MaxInt
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if i == 0 {
				// first row
				res[i][j] = matrix[i][j]
				continue
			}

			topRes := res[i-1][j]
			leftRes := math.MaxInt
			rightRes := math.MaxInt

			if j == 0 {
				// left side
				rightRes = res[i-1][j+1]
			} else if j == len(matrix[i])-1 {
				// right side
				leftRes = res[i-1][j-1]
			} else {
				// else
				rightRes = res[i-1][j+1]
				leftRes = res[i-1][j-1]
			}
			res[i][j] = min([]int{topRes, leftRes, rightRes}) + matrix[i][j]

			if i == len(matrix)-1 && ret > res[i][j] {
				ret = res[i][j]
			}
		}
	}
	return ret
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
