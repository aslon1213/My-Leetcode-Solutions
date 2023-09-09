package main

import "fmt"

func main() {

	fmt.Println(combinationSum4([]int{2, 1, 3}, 32))

}

func combinationSum4(nums []int, target int) int {

	ways := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] > target {
			continue
		}
		if nums[i] == target {
			ways++
			continue
		}
		// fmt.Println(nums[i])
		ways += combinationSum4(nums, target-nums[i])

	}

	return ways

}
