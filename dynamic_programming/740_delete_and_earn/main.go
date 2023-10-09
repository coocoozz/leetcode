package main

import (
	"fmt"
	"leetcode/utils"
)

func main() {
	testCases := []struct {
		nums []int
	}{
		{[]int{3, 4, 2}},
		{[]int{2, 2, 3, 3, 3, 4}},
		{[]int{4, 3, 2, 1}},
		{[]int{2, 2, 3, 3, 3, 4, 5}},
		{[]int{1, 1, 1, 2, 4, 5, 5, 5, 6}},
	}

	for _, tc := range testCases {
		fmt.Printf("* Case %+v\n", tc.nums)

		res := 0
		elapsed := utils.ElapsedTime(func() {
			res = deleteAndEarnRecursive(tc.nums)
		})
		fmt.Printf("  - res: %d, elapsed: %v\n", res, elapsed)

		res = 0
		elapsed = utils.ElapsedTime(func() {
			res = deleteAndEarnRecursiveWithMem(tc.nums)
		})
		fmt.Printf("  - resWithMem: %d, elapsed: %v\n", res, elapsed)

		res = 0
		elapsed = utils.ElapsedTime(func() {
			res = deleteAndEarnWithBT(tc.nums)
		})
		fmt.Printf("  - resWithBT: %d, elapsed: %v\n", res, elapsed)
	}
}

func makePointMap(nums []int) (int, map[int]int) {
	max := 0
	pointMap := make(map[int]int)
	for _, num := range nums {
		pointMap[num] += num
		max = utils.Max(max, num)
	}
	return max, pointMap
}

/** Recurrence Relation
  f(n): 가장 높은 point n부터 시작하여 모든 case에 대한 최대 points for delete and earn opration
    - 0 => n < 0
	- max(f(n-2) + point(n), f(n-1)) => n > 0
**/

func deleteAndEarnRecursive(nums []int) int {
	max, pointMap := makePointMap(nums)
	return recursive(max, pointMap)
}

func recursive(max int, pointMap map[int]int) int {
	if max < 1 {
		return 0
	}
	selected := recursive(max-2, pointMap) + pointMap[max]
	notSelected := recursive(max-1, pointMap)
	return utils.Max(selected, notSelected)
}

func deleteAndEarnRecursiveWithMem(nums []int) int {
	max, pointMap := makePointMap(nums)
	mem := make(map[int]int)
	return recursiveWithMem(max, pointMap, mem)
}

func recursiveWithMem(max int, pointMap map[int]int, mem map[int]int) int {
	if max < 1 {
		return 0
	} else if val, exist := mem[max]; exist {
		return val
	}

	selected := recursiveWithMem(max-2, pointMap, mem) + pointMap[max]
	notSelected := recursiveWithMem(max-1, pointMap, mem)
	mem[max] = utils.Max(selected, notSelected)
	return mem[max]
}

func deleteAndEarnWithBT(nums []int) int {
	max, pointMap := makePointMap(nums)
	res := make([]int, max+1)
	res[0] = 0
	res[1] = 0
	if val, exist := pointMap[1]; exist {
		res[1] = val
	}

	for i := 2; i <= max; i++ {
		selected := res[i-2] + pointMap[i]
		notSelected := res[i-1]
		res[i] = utils.Max(selected, notSelected)
	}
	fmt.Printf("%+v\n", res)
	return res[max]
}
