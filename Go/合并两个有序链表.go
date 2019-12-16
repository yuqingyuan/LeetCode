package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 递归
// func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
// 	if l1 == nil {
// 		return l2
// 	}
// 	if l2 == nil {
// 		return l1
// 	}
// 	if l1.Val < l2.Val {
// 		return &ListNode{l1.Val, mergeTwoLists(l1.Next, l2)}
// 	} else {
// 		return &ListNode{l2.Val, mergeTwoLists(l1, l2.Next)}
// 	}
// }

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := l1
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	} else {
		for {
			if l2 == nil {
				break
			}
			if l2.Val < l1.Val {
				node := &ListNode{l2.Val, l1}
				head = node
				l1 = head
			} else {
				for {
					if l1.Next == nil {
						l1.Next = &ListNode{l2.Val, nil}
						l1 = l1.Next
						break
					}
					if l1.Next.Val > l2.Val {
						node := &ListNode{l2.Val, l1.Next}
						l1.Next = node
						break
					} else {
						l1 = l1.Next
					}
				}
			}
			l2 = l2.Next
		}
	}
	return head
}

func main() {
	h1 := &ListNode{2, nil}
	h2 := &ListNode{1, nil}
	fmt.Println(mergeTwoLists(h1, h2))
}

