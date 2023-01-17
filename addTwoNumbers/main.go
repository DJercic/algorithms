package addTwoNumbers

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	node := ListNode{}
	head := &node
	carry := 0
	for l1.Next != nil || l2.Next != nil {
		node.Val = (l1.Val + l2.Val + carry) % 10
		carry = int(l1.Val + l2.Val + carry/10)
	}
	if carry > 0 {
		node.Next = &ListNode{Val: carry}
	}
	return head
}
