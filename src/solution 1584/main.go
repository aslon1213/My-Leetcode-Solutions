package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minCostConnectPoints([][]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}}))
	fmt.Println(minCostConnectPoints([][]int{{3, 12}, {-2, 5}, {-4, 1}}))

}

type Points [][]int

func (p Points) Len() int {
	return len(p)
}

func (p Points) Less(i, j int) bool {
	return p[i][0] < p[j][0]
}

func (p Points) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func minCostConnectPoints(points [][]int) int {
	n := 0

	// sort by x
	sort.Sort(Points(points))
	fmt.Println(points)
	// sort by y
	// left_bottom := points[0]

	for i := 0; i < len(points); i++ {

		// maximum int

		min_dist := 1000000000
		p := 0
		for j := i + 1; j < len(points); j++ {
			dist := dist(points[i], points[j])
			if dist < min_dist {
				min_dist = dist
				p = j
			}
		}
		points = append(points[:p], points[p+1:]...)
		n += min_dist
	}

	return n
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func dist(a, b []int) int {
	return abs(a[0]-b[0]) + abs(a[1]-b[1])
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
