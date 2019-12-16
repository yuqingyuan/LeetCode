package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var flag int
	var head *ListNode
	var curr *ListNode = &ListNode{0, nil}
	head = curr
	for {
		var x, y int
		if l1 != nil {
			x = l1.Val
		} else {
			x = 0
		}
		if l2 != nil {
			y = l2.Val
		} else {
			y = 0
		}
		sum := x + y + flag
		flag = sum / 10
		curr.Next = &ListNode{sum % 10, nil}
		curr = curr.Next

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
		if l1 == nil && l2 == nil {
			break
		}
	}
	if flag == 1 {
		curr.Next = &ListNode{1, nil}
	}
	return head.Next
}

func main() {
	l1 := ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 := ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	addTwoNumbers(&l1, &l2)
}

