package addTwoNumbers

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	var head *ListNode
	var previous *ListNode
	n1 := l1
	n2 := l2
	for n1 != nil || n2 != nil {
		node := ListNode{}
		node.Val = (sum(n1, n2) + carry) % 10
		if head == nil {
			head = &node
		}
		carry = int((sum(n1, n2) + carry) / 10)
		if n1 != nil {
			n1 = n1.Next
		}
		if n2 != nil {
			n2 = n2.Next
		}
		if previous != nil {
			previous.Next = &node
		}
		previous = &node
	}
	if carry > 0 {
		previous.Next = &ListNode{Val: carry}
	}
	return head
}

func sum(l1 *ListNode, l2 *ListNode) int {
	if l1 != nil && l2 != nil {
		return l1.Val + l2.Val
	} else if l1 != nil {
		return l1.Val
	} else {
		return l2.Val
	}
}
