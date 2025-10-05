package main

import "fmt"

func main() {

	cases := [][]int{
		{1, 8, 6, 2, 5, 4, 8, 3, 7},
		{1, 1},
	}
	answers := []int{
		49,
		1,
	}
	for i, _case := range cases {
		got := maxArea(_case)
		if got != answers[i] {
			fmt.Printf("maxArea(%v) = %d, want %d\n", _case, got, answers[i])
		} else {
			fmt.Printf("maxArea(%v) = %d, want %d\n", _case, got, answers[i])
		}
	}
}

func maxArea(height []int) int {

	max_area := 0
	left_point := 0
	right_point := len(height) - 1
	for left_point < right_point {
		left_val, right_val := height[left_point], height[right_point]
		area := (right_point - left_point) * min(left_val, right_val)
		if area > max_area {
			max_area = area
		}
		if left_val < right_val {
			left_point++
		} else {
			right_point--
		}
	}
	return max_area
}
