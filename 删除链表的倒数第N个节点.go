package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	p := head
	q := head
	for i := 0; i < n; i++ {
		p = p.Next
	}
	if p == nil {
		head = head.Next
		return head
	} else {
		for {
			p = p.Next
			if p == nil {
				break
			}
			q = q.Next
		}
		q.Next = q.Next.Next
	}

	return head
}

func main() {
	head := &ListNode{1, &ListNode{2, nil}}
	fmt.Println(removeNthFromEnd(head, 2))
}
