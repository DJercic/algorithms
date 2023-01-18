package addTwoNumbers

import "math"

/**
* Definition for singly-linked list.
**/

type ListNode struct {
	Val  int
	Next *ListNode
}

// add function length to ListNode struct
func (l *ListNode) length() int {
	if l == nil {
		return 0
	}
	return 1 + l.Next.length()
}

// add function toNumber to ListNode struct
func (l *ListNode) toNumber() int {
	result := 0
	node := l
	for i := 0; i < l.length(); i++ {
		result += node.Val * int(math.Pow10(i))
		node = node.Next
	}
	return result
}
