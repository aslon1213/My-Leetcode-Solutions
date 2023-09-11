package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
)

type BinaryTreeNode struct {
	Val   int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

func (b *BinaryTreeNode) AddNode(val int) {
	if b == nil || b.Val == 0 {
		b.Val = val
		return
	}
	if b.Val == val {
		return
	}

	if val < b.Val {
		if b.Left == nil {
			b.Left = &BinaryTreeNode{Val: val, Left: nil, Right: nil}
			return
		}
		b.Left.AddNode(val)
		return
	}
	if b.Right == nil {
		b.Right = &BinaryTreeNode{Val: val, Left: nil, Right: nil}
		return
	}
	b.Right.AddNode(val)

}
func (b *BinaryTreeNode) String() {
	if b == nil {
		return
	}
	fmt.Println(b.Val)
	b.Left.String()
	b.Right.String()
}

func (b *BinaryTreeNode) InOrderTraversal() []int {
	if b == nil {
		return []int{}
	}
	return append(append(b.Left.InOrderTraversal(), b.Val), b.Right.InOrderTraversal()...)
}

func main() {

	// open file
	dir := os.Getenv("PWD")
	file, err := os.OpenFile(path.Join(dir, "README.md"), os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file_contents := []string{}
	// read file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		file_contents = append(file_contents, scanner.Text())
	}
	solutions := file_contents[4:]
	solutions_map := map[int]string{}
	var b_tree BinaryTreeNode
	for i := 0; i < len(solutions); i++ {
		num_s := strings.Split(solutions[i], "|")
		fmt.Println(num_s)
		num, err := strconv.Atoi(strings.TrimSpace(num_s[1]))
		if err != nil {
			panic(err)
		}
		solutions_map[num] = solutions[i]
		b_tree.AddNode(num)
	}
	traversed := b_tree.InOrderTraversal()
	for i := 0; i < len(traversed); i++ {
		file_contents[i+4] = solutions_map[traversed[i]]
	}
	// write file
	file.Truncate(0)
	file.Seek(0, 0)
	for _, line := range file_contents {
		fmt.Fprintln(file, line)
	}

}
