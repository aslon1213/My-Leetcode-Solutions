// a program to add a new solution to Readme.md
package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	//get the current directory
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}

	// get the solution number
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter solution number: ")

	text, _, _ := reader.ReadLine()
	solutionNumber, err := strconv.Atoi(string(text))
	if err != nil {
		fmt.Println(err)
		return
	}
	// construct a link
	link := fmt.Sprintf("https://github.com/aslon1213/My-Leetcode-Solutions/tree/main/src/solution%%20%d", solutionNumber)
	// construct a new line
	newLine := fmt.Sprintf("| %d | %s |%s |\n", solutionNumber, "go", link)
	// write the new line to the readme file
	f, err := os.OpenFile(filepath.Join(dir, "Readme.md"), os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	if _, err = f.WriteString(newLine); err != nil {
		fmt.Println(err)
		return
	}

}
