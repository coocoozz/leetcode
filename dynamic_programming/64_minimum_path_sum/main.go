package main

import (
	"fmt"
	"leetcode/utils"
)

func main() {
	testCases := []struct {
		grid     [][]int
		expected int
	}{
		{
			[][]int{
				{1, 3, 1},
				{1, 5, 1},
				{4, 2, 1},
			},
			7,
		},
		{
			[][]int{
				{1, 2, 3},
				{4, 5, 6},
			},
			12,
		},
	}

	for i, tc := range testCases {
		fmt.Printf("* Case %d => %d\n", i, tc.expected)
		var res int

		elapsed := utils.ElapsedTime(func() {
			res = minPathSumRecursive(tc.grid)
		})
		fmt.Printf("  - recursive: %d, elapsed: %v\n", res, elapsed)

		elapsed = utils.ElapsedTime(func() {
			res = minPathSumMem(tc.grid)
		})
		fmt.Printf("  - mem: %d, elapsed: %v\n", res, elapsed)

		elapsed = utils.ElapsedTime(func() {
			res = minPathSumBT(tc.grid)
		})
		fmt.Printf("  - BT: %d, elapsed: %v\n", res, elapsed)
	}
}

/** Recurrence Relation

  cost(i, j) = (i, j)에서의 cost
  f(i, j) = (i, j)에 다다르기까지 min cost
    - MaxInt => i < 0 or j < 0
	- cost(0, 0) => i == 0 and j == 0
	- f(i, j) = min(f(i-1, j), f(i, j-1)) + cost(i, j)

**/

const MAX = 201

func minPathSumRecursive(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	return recursive(m-1, n-1, grid)
}

func recursive(m, n int, grid [][]int) int {
	if m == 0 && n == 0 {
		return grid[0][0]
	} else if m < 0 || n < 0 {
		return MAX
	}

	topRes := recursive(m-1, n, grid)
	leftRes := recursive(m, n-1, grid)
	return utils.Min(topRes, leftRes) + grid[m][n]
}

func minPathSumMem(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	mem := make([][]int, m)
	for i := 0; i < m; i++ {
		mem[i] = make([]int, n)
	}
	return recursiveMem(m-1, n-1, grid, mem)
}

func recursiveMem(m, n int, grid [][]int, mem [][]int) int {
	if m == 0 && n == 0 {
		return grid[0][0]
	} else if m < 0 || n < 0 {
		return MAX
	} else if mem[m][n] != 0 {
		return mem[m][n]
	}

	topRes := recursiveMem(m-1, n, grid, mem)
	leftRes := recursiveMem(m, n-1, grid, mem)
	mem[m][n] = utils.Min(topRes, leftRes) + grid[m][n]
	return mem[m][n]
}

func minPathSumBT(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				res[i][j] = grid[0][0]
			} else if i == 0 && j > 0 {
				res[i][j] = res[0][j-1] + grid[0][j]
			} else if j == 0 && i > 0 {
				res[i][j] = res[i-1][0] + grid[i][0]
			} else {
				res[i][j] = utils.Min(res[i-1][j], res[i][j-1]) + grid[i][j]
			}
		}
	}
	return res[m-1][n-1]
}
