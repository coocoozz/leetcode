package main

import (
	"fmt"
	"leetcode/utils"
)

func main() {
	testCases := []struct {
		triangle [][]int
		expected int
	}{
		{
			[][]int{
				{2},
				{3, 4},
				{6, 5, 7},
				{4, 1, 8, 3},
			},
			11,
		},
		{
			[][]int{
				{-10},
			},
			-10,
		},
	}

	for i, tc := range testCases {
		fmt.Printf("* Case[%d] => %d\n", i, tc.expected)
		var res int

		elapsed := utils.ElapsedTime(func() {
			res = minimumTotalRecursive(tc.triangle)
		})
		fmt.Printf("  - recursive => res: %d, elapsed: %v\n", res, elapsed)

		elapsed = utils.ElapsedTime(func() {
			res = minimumTotalBT(tc.triangle)
		})
		fmt.Printf("  - BT => res: %d, elapsed: %v\n", res, elapsed)
	}

}

/** Recurrence Relations

  cose(i, j) = (i, j) 에서의 cost
  f(i ,j) => triangle bottom에서 부터 (i, j) 까지 도달하기 까지 minimum path sum
    - cost(i, j) => i == triangle bottom
	- 0 => i > triangle bottom || j > length of triangle bottom
	- f(i, j) = min(f(i+1, j), f(i+1, j+1)) + cost(i, j)

**/

func minimumTotalRecursive(triangle [][]int) int {
	return recursive(0, 0, triangle)
}

func recursive(i, j int, triangle [][]int) int {
	bottomIdx := len(triangle) - 1
	if i > bottomIdx {
		return 0
	} else if i == bottomIdx {
		return triangle[i][j]
	}

	downRes := recursive(i+1, j, triangle)
	rightRes := recursive(i+1, j+1, triangle)
	return utils.Min(downRes, rightRes) + triangle[i][j]
}

func minimumTotalBT(triangle [][]int) int {
	res := make([][]int, len(triangle))
	for i := 0; i < len(triangle); i++ {
		res[i] = make([]int, len(triangle[i]))
	}

	for i := len(triangle) - 1; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			if i == len(triangle)-1 {
				res[i][j] = triangle[i][j]
			} else {
				res[i][j] = utils.Min(res[i+1][j], res[i+1][j+1]) + triangle[i][j]
			}
		}
	}

	return res[0][0]
}
