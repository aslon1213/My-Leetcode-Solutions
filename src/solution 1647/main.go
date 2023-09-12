package main

import (
	"fmt"
)

func main() {
	fmt.Println(minDeletions("aab"))
	fmt.Println(minDeletions("aaabbbcc"))
	fmt.Println(minDeletions("ceabaacb"))
}
func minDeletions(s string) int {
	var count [26]int
	used := make(map[int]int)

	for _, ch := range s {
		count[ch-'a']++
	}

	var ans int
	for i := 0; i < 26; i++ {
		check := count[i]
		for used[check] != 0 && check != 0 {
			check--
		}
		ans += count[i] - check
		used[check] = 1
	}

	return ans
}
