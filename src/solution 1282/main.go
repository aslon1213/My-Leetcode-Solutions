package main

import "fmt"

func main() {

	fmt.Println(groupThePeople([]int{3, 3, 3, 3, 3, 1, 3}))

}

func groupThePeople(groupSizes []int) [][]int {

	group_sizes := make(map[int][]int)
	for i, num := range groupSizes {
		if _, ok := group_sizes[num]; !ok {
			group_sizes[num] = []int{}
		}
		group_sizes[num] = append(group_sizes[num], i)
	}
	fmt.Println(group_sizes)
	output := [][]int{}
	for k, v := range group_sizes {
		for i := 0; i < len(v)/k; i++ {
			output = append(output, v[i*k:(i+1)*k])
		}
	}

	return output
}
