package main

import (
	"fmt"
	"leetcode/utils"
)

func main() {
	testCases := []struct {
		obstacleGrid [][]int
		expected     int
	}{
		{
			[][]int{
				{0, 0, 0},
				{0, 1, 0},
				{0, 0, 0},
			},
			2,
		},
		{
			[][]int{
				{0, 1},
				{0, 0},
			},
			1,
		},
		{
			[][]int{
				{1, 0},
			},
			0,
		},
	}

	for i, tc := range testCases {
		fmt.Printf("* Case %d => %d\n", i, tc.expected)
		var res int

		elapsed := utils.ElapsedTime(func() {
			res = uniquePathsWithObstaclesRecursive(tc.obstacleGrid)
		})
		fmt.Printf("  - recursive: %d, elapsed: %v\n", res, elapsed)

		elapsed = utils.ElapsedTime(func() {
			res = uniquePathsWithObstaclesBT(tc.obstacleGrid)
		})
		fmt.Printf("  - BT: %d, elapsed: %v\n", res, elapsed)
	}
}

/** Recurrence Relations

  f(i, j) = 장애물을 피해서 (i, j) 까지 올 수 있는 모든 경우의 수
    - 0 => (i, j) is obstacle and i < 0 or j < 0
    - 1 => i == 0 and j == 0 and not obstacle
	- f(i, j)
	  - f(i-1, j) + f(i, j-1) => f(i-1, j) and f(i, j-1) are not obstacle
	  - f(i-1, j) => f(i-1, j) is not obstacle and f(i, j-1) is an obstacle
	  - f(i, j-1) => f(i, j-1) is not obstacle and f(i-1, j) is an obstacle

**/

func uniquePathsWithObstaclesRecursive(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	return recursive(m-1, n-1, obstacleGrid)
}

func recursive(m, n int, obstacleGrid [][]int) int {
	if m < 0 || n < 0 {
		return 0
	} else if obstacleGrid[m][n] == 1 {
		return 0
	} else if m == 0 && n == 0 {
		return 1
	}
	return recursive(m-1, n, obstacleGrid) + recursive(m, n-1, obstacleGrid)
}

func uniquePathsWithObstaclesBT(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				res[i][j] = 0
			} else if i == 0 && j > 0 {
				res[i][j] = res[i][j-1]
			} else if j == 0 && i > 0 {
				res[i][j] = res[i-1][j]
			} else if i == 0 && j == 0 {
				res[i][j] = 1
			} else {
				res[i][j] = res[i-1][j] + res[i][j-1]
			}
		}
	}
	return res[m-1][n-1]
}
