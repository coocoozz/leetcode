package main

import (
	"fmt"
	"time"
)

func main() {
	n := 44

	now := time.Now()
	res1 := climbStairsWithMem(n)
	res1Elapsed := time.Now().Sub(now)

	now = time.Now()
	res2 := climbStairsWithMem(n)
	res2Elapsed := time.Now().Sub(now)

	now = time.Now()
	res3 := climbStairsWithBT(n)
	res3Elapsed := time.Now().Sub(now)

	fmt.Printf("n: %d, res1: %d, elapsed: %v\n", n, res1, res1Elapsed)
	fmt.Printf("n: %d, res2: %d, elapsed: %v\n", n, res2, res2Elapsed)
	fmt.Printf("n: %d, res3: %d, elapsed: %v\n", n, res3, res3Elapsed)
}

func climbStairs(n int) int {
	if n == 0 {
		return 1
	} else if n < 0 {
		return 0
	}

	return climbStairs(n-1) + climbStairs(n-2)
}

var mem map[int]int = make(map[int]int)

func climbStairsWithMem(n int) int {
	if val, exist := mem[n]; exist {
		return val
	} else if n == 0 {
		return 1
	} else if n < 0 {
		return 0
	}

	mem[n] = climbStairs(n-1) + climbStairs(n-2)
	return mem[n]
}

func climbStairsWithBT(n int) int {
	res := make([]int, n+1)
	res[0] = 1

	for i := 1; i <= n; i++ {
		oneStepRes := i - 1
		twoStepRes := i - 2

		if oneStepRes == 0 {
			oneStepRes = 1
		} else if oneStepRes < 0 {
			oneStepRes = 0
		} else {
			oneStepRes = res[oneStepRes]
		}

		if twoStepRes == 0 {
			twoStepRes = 1
		} else if twoStepRes < 0 {
			twoStepRes = 0
		} else {
			twoStepRes = res[twoStepRes]
		}

		res[i] = oneStepRes + twoStepRes
	}
	return res[n]
}
