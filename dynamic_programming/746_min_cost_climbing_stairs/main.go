package main

import (
	"fmt"
	"leetcode/utils"
	"math"
)

func main() {
	cost := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}

	res1 := 0
	elapsed1 := utils.ElapsedTime(func() {
		res1 = minCostClimbingStairs(len(cost), cost)
	})
	fmt.Printf("res1: %d, elapsed: %v\n", res1, elapsed1)

	res2 := 0
	elapsed2 := utils.ElapsedTime(func() {
		res2 = minCostClimbingStairsWithBT(cost)
	})
	fmt.Printf("res2: %d, elapsed: %v\n", res2, elapsed2)
}

/** Recurrence Relation

cost(step) = 현재 step에서의 cost
f(step) = 현재 step 까지의 누적되 cost의 최소값

f(step)
	- 0 => if step is length of cost
	- cost(step) => if step is 0 or 1
	- min(f(step - 1), f(step - 2)) + cost(step) => step is greater than 2

*/

func minCostClimbingStairs(step int, cost []int) int {
	if step == 0 {
		return cost[0]
	} else if step == 1 {
		return cost[1]
	}

	oneStepBeforeCost := minCostClimbingStairs(step-1, cost)
	twoStepBeforeCost := minCostClimbingStairs(step-2, cost)
	curCost := 0
	if step < len(cost) {
		curCost = cost[step]
	}

	return int(math.Min(float64(oneStepBeforeCost), float64(twoStepBeforeCost))) + curCost
}

func minCostClimbingStairsWithBT(cost []int) int {
	totalSteps := len(cost) + 1

	minCostList := make([]int, totalSteps)
	minCostList[0] = cost[0]
	minCostList[1] = cost[1]

	for i := 2; i < totalSteps; i++ {
		curStepCost := 0
		if i < len(cost) {
			curStepCost = cost[i]
		}

		minCostList[i] = int(math.Min(float64(minCostList[i-1]), float64(minCostList[i-2]))) + curStepCost
	}
	return minCostList[len(cost)]
}
