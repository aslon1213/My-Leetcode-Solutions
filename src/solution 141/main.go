package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return &ListNode{Val: val, Next: nil}
	}
	cur := head
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = &ListNode{Val: val, Next: nil}
	return head
}

func (l *ListNode) String() {
	cur := l
	for cur != nil {
		fmt.Printf("%d -> ", cur.Val)
		cur = cur.Next
	}
	fmt.Println()
}

func main() {
	head := &ListNode{Val: 1, Next: nil}
	addNode(head, 2)
	addNode(head, 3)
	addNode(head, 4)
	addNode(head, 5)
	node := &ListNode{Val: 6, Next: nil}
	node.Next = head
	head.Next.Next.Next.Next.Next = node
	fmt.Println(hasCycle(head))
}

func hasCycle(head *ListNode) bool {

	cur := head
	nodes := []*ListNode{}
	for {
		nodes = append(nodes, cur)
		// check if contains
		for i := 0; i < len(nodes); i++ {
			if nodes[i] == cur.Next {
				return true
			}
		}
		if cur.Next == nil {
			return false
		}
		cur = cur.Next
	}

}
