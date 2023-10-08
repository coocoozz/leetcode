package main

import (
	"fmt"
	"leetcode/utils"
)

func main() {
	testCases := []struct {
		nums []int
	}{
		{[]int{1, 2, 3, 1}},
		{[]int{2, 7, 9, 3, 1}},
		{[]int{2, 1, 1, 2}},
	}

	for _, tc := range testCases {
		fmt.Printf("* Case %+v\n", tc.nums)
		var res int

		elapsed := utils.ElapsedTime(func() {
			res = robRecursive(len(tc.nums)-1, tc.nums)
		})
		fmt.Printf("  - resRecursive: %d, elapsed: %v\n", res, elapsed)

		elapsed = utils.ElapsedTime(func() {
			res = robRecursiveWithMem(len(tc.nums)-1, tc.nums)
		})
		fmt.Printf("  - resRecursiveWithMem: %d, elapsed: %v\n", res, elapsed)

		elapsed = utils.ElapsedTime(func() {
			res = robWithBT(tc.nums)
		})
		fmt.Printf("  - resWithBT: %d, elapsed: %v\n", res, elapsed)

	}
}

/** Recurrence Relation

rob(i): i 번째 house까지 최대로 털 수 있는 money
 - 0 => i < 0
 - Max(rob(i-2) + currentHouseMoney, rob(i-1)) => i >= 0

*/

func robRecursive(house int, nums []int) int {
	if house < 0 {
		return 0
	}

	selected := robRecursive(house-2, nums) + nums[house]
	notSelected := robRecursive(house-1, nums)
	return utils.Max(selected, notSelected)
}

var mem map[int]int = make(map[int]int)

func robRecursiveWithMem(house int, nums []int) int {
	if val, exist := mem[house]; exist {
		return val
	} else if house < 0 {
		return 0
	}

	selected := robRecursive(house-2, nums) + nums[house]
	notSelected := robRecursive(house-1, nums)
	mem[house] = utils.Max(selected, notSelected)
	return mem[house]
}

func robWithBT(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}

	res := make([]int, len(nums))
	res[0] = nums[0]
	res[1] = utils.Max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		selected := res[i-2] + nums[i]
		notSelected := res[i-1]
		res[i] = utils.Max(selected, notSelected)
	}
	return res[len(nums)-1]
}
