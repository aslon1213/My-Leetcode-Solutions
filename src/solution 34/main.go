package main

import "fmt"

type Case struct {
	nums   []int
	target int
	want   []int
}

func main() {

	cases := []Case{
		{nums: []int{5, 7, 7, 8, 8, 10}, target: 8, want: []int{3, 4}},
		{nums: []int{5, 7, 7, 8, 8, 10}, target: 6, want: []int{-1, -1}},
		{nums: []int{1}, target: 1, want: []int{0, 0}},
	}
	for _, c := range cases {
		got := searchRange(c.nums, c.target)

		if got[0] != c.want[0] || got[1] != c.want[1] {
			fmt.Printf("Failed! searchRange(%v, %d) = %v, want %v\n", c.nums, c.target, got, c.want)
		} else {
			fmt.Printf("Passed! searchRange(%v, %d) = %v\n", c.nums, c.target, got)
		}
	}
}

func searchRange(nums []int, target int) []int {
	left, right := -1, -1
	for i, x := range nums {
		if x == target {
			if left == -1 {
				left = i
			}
			right = i
		}
	}
	return []int{left, right}
}
