package main

import (
	"fmt"
	"leetcode/utils"
)

func main() {
	testCases := []struct {
		target []int
	}{
		{[]int{3, 7}},
		{[]int{3, 2}},
	}

	for _, tc := range testCases {
		fmt.Printf("* Case: %d, %d\n", tc.target[0], tc.target[1])
		var res int

		elapsed := utils.ElapsedTime(func() {
			res = uniquePathsDPRecursive(tc.target[0], tc.target[1])
		})
		fmt.Printf("  - recursive: %d, elapsed: %v\n", res, elapsed)

		elapsed = utils.ElapsedTime(func() {
			res = uniquePathsDPMem(tc.target[0], tc.target[1])
		})
		fmt.Printf("  - mem: %d, elapsed: %v\n", res, elapsed)

		elapsed = utils.ElapsedTime(func() {
			res = uniquePathsDP(tc.target[0], tc.target[1])
		})
		fmt.Printf("  - DP: %d, elapsed: %v\n", res, elapsed)

	}
}

/** Recurrence Relations

  f(i, j) = (i, j)까지 도달할 수 있는 모든 경우의 수
    - f(0, 0) = 0
    - f(0, ...) = 1
	- f(..., 0) = 1
	- f(i, j) = f(i-1, j) + f(i, j-1)

**/

func uniquePathsDPRecursive(m int, n int) int {
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
	}
	recursive(m-1, n-1, res)
	return res[m-1][n-1]
}

func recursive(i, j int, res [][]int) int {
	if i == 0 && j == 0 {
		return 0
	} else if i == 0 && j > 0 {
		return 1
	} else if j == 0 && i > 0 {
		return 1
	}

	res[i][j] = recursive(i-1, j, res) + recursive(i, j-1, res)
	return res[i][j]
}

func uniquePathsDPMem(m int, n int) int {
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
	}
	recursiveMem(m-1, n-1, res)
	return res[m-1][n-1]
}

func recursiveMem(i, j int, res [][]int) int {
	if i == 0 && j == 0 {
		return 0
	} else if i == 0 && j > 0 {
		return 1
	} else if j == 0 && i > 0 {
		return 1
	} else if res[i][j] != 0 {
		return res[i][j]
	}

	res[i][j] = recursiveMem(i-1, j, res) + recursiveMem(i, j-1, res)
	return res[i][j]
}

func uniquePathsDP(m int, n int) int {
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if (i == 0 && j == 0) || (i == 0 && j > 0) || (j == 0 && i > 0) {
				res[i][j] = 1
			} else {
				res[i][j] = res[i-1][j] + res[i][j-1]
			}
		}
	}
	return res[m-1][n-1]
}
