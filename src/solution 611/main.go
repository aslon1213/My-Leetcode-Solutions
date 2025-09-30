package main

import "fmt"

func main() {
	cases := [][]int{
		[]int{1, 2, 3, 4, 5},
		[]int{5},
	}
	for _, _case := range cases {
		got := triangularSum(_case)
		fmt.Println(got)
	}

}

func triangularSum(nums []int) int {

	for len(nums) > 1 {
		newNums := make([]int, len(nums)-1)
		for i := 0; i < len(nums)-1; i++ {
			newNums[i] = (nums[i] + nums[i+1]) % 10
		}
		nums = newNums
	}

	return nums[0]
}
