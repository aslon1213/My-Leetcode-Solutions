package main

import (
	"fmt"
	"strconv"
)



type Case struct {
	s string
	want bool
}

func main() {
	cases := []Case{
		{s: "3902", want: true},
		{s: "34789", want: false},
	}

	for _, c := range cases {
		got := hasSameDigits(c.s)
		if got != c.want {
			fmt.Printf("Test failed! hasSameDigits(%s) = %t, want %t\n", c.s, got, c.want)
		}else {
			fmt.Printf("Test passed! hasSameDigits(%s) = %t\n", c.s, got)
		}
	}

}


func hasSameDigits(s string) bool {
	// fmt.Println(s)
    if len(s) == 2 && s[0] == s[1] {
        return true
    } else if len(s) == 2 {
        return false
    }

	var new_s string
	for i :=0; i < len(s) - 1; i++ {
		left, _ := strconv.Atoi(string(s[i]))
		right, _ := strconv.Atoi(string(s[i+1]))
		new_s += strconv.Itoa((left + right ) % 10)

	}
	return hasSameDigits(new_s)
    
}