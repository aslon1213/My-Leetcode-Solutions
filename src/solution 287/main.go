package main

import "fmt"

func main() {
	fmt.Println(findDuplicate([]int{1, 3, 4, 2, 2}))
	fmt.Println(findDuplicate([]int{3, 1, 3, 4, 2}))

}

func findDuplicate(nums []int) int {
	nums_set := make([]int, len(nums))
	for _, v := range nums {
		nums_set[v]++
		if nums_set[v] > 1 {
			return v
		}
	}
	return 0
}
