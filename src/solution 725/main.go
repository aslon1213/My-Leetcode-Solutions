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
	addNode(head, 6)
	addNode(head, 7)
	addNode(head, 8)
	addNode(head, 9)
	addNode(head, 10)
	head.String()
	lists := splitListToParts(head, 3)
	for _, list := range lists {
		list.String()
	}

}

func splitListToParts(head *ListNode, k int) []*ListNode {

	// get length of list
	length := 0
	cur := head
	for {
		if cur == nil {
			break
		}
		length++
		cur = cur.Next
	}
	fmt.Printf("LENGtH: %d\n", length)

	r := length % k
	q := length / k
	lists := make([]*ListNode, k)
	cur = head
	for i := 0; i < r; i++ {
		list_head := head
		for j := 0; j < q; j++ {
			cur = cur.Next
		}
		head = cur.Next
		cur.Next = nil
		cur = head

		lists[i] = list_head
	}
	for i := r; i < k; i++ {
		if cur == nil {
			lists[i] = nil
			continue
		}
		list_head := head
		for j := 0; j < q-1; j++ {
			cur = cur.Next
		}
		head = cur.Next
		cur.Next = nil
		cur = head
		lists[i] = list_head
	}

	return lists

}
