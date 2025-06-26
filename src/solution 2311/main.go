package main

func main() {

	type Case struct {
		s    string
		k    int
		want int
	}
	cases := []Case{
		{s: "1001010", k: 5, want: 5},
		{s: "00101001", k: 1, want: 6},
	}
	for _, c := range cases {
		// fmt.Printf("\nTesting case: s=%s, k=%d\n", c.s, c.k)
		got := longestSubsequence(c.s, c.k)
		if got != c.want {
			// fmt.Printf("longestSubsequence(%s, %d) = %d, want %d\n", c.s, c.k, got, c.want)
		} else {
			// fmt.Printf("Test passed! Got %d\n", got)
		}
	}

}
func longestSubsequence(s string, k int) int {
	n := len(s)
	count := 0
	power := 0
	val := int64(0)

	for i := n - 1; i >= 0; i-- {
		if s[i] == '0' {
			count++
		} else {
			if power < 32 {
				val += 1 << power
				if val <= int64(k) {
					count++
				}
			}
		}
		power++
	}
	return count
}
