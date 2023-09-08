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

	node := &ListNode{Val: 1, Next: nil}
	addNode(node, 2)
	addNode(node, 3)
	addNode(node, 4)
	addNode(node, 5)
	reverseBetween(node, 1, 3).String()

}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head.Next == nil || left == right {

		return head
	}
	if head.Next.Next == nil {

		swap := head.Next
		swap.Next = head
		head.Next = nil

		return swap
	}

	pre_left := &ListNode{Val: 0, Next: head}
	cur := head
	for i := 1; i < right; i++ {
		if left == 1 {
			left_node := &ListNode{Val: 0, Next: head}
			left_node.Val = cur.Next.Val
			cur.Next = cur.Next.Next
			head = left_node
			continue
		}
		if i < left {
			pre_left = cur
			cur = cur.Next
			continue
		}
		swap := cur.Next
		swap_2 := pre_left.Next

		pre_left.Next = swap
		cur.Next = swap.Next
		swap.Next = swap_2
	}
	return head

}
