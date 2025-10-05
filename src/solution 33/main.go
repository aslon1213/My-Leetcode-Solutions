package main

import "fmt"

type Case struct {
	nums   []int
	target int
	want   int
}

func main() {
	cases := []Case{
		{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 0, want: 4},
		{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 3, want: -1},
		{nums: []int{1}, target: 0, want: -1},
	}
	for _, c := range cases {
		got := search(c.nums, c.target)
		if got != c.want {
			fmt.Printf("Test failed! search(%v, %d) = %d, want %d\n", c.nums, c.target, got, c.want)
		} else {
			fmt.Printf("Test passed! search(%v, %d) = %d\n", c.nums, c.target, got)
		}
	}
}

func search(nums []int, target int) int {
	for i, x := range nums {
		if x == target {
			return i
		}
	}
	return -1
}
