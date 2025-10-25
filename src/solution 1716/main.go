package main

import "fmt"


type Case struct {
	n int
	want int
}

func main() {
	cases := []Case{
		{n: 4, want: 10},
		{n: 10, want: 37},
		{n: 20, want: 96},
	}
	for _, c := range cases {
		got := totalMoney(c.n)
		if got != c.want {
			fmt.Printf("totalMoney(%d) = %d, want %d\n", c.n, got, c.want)
		}else {
			fmt.Printf("totalMoney(%d) = %d, want %d\n", c.n, got, c.want)
		}
	}
}



func totalMoney(n int) int {
    money := 0
	for i := 1; i <= n; i++ {
		week_day := i % 7
		if week_day == 0 {
			week_day = 6
		}
		money += week_day + (i / 7)
	}
	return money
}