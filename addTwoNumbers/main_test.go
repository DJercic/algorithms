package addTwoNumbers

import "testing"

func TestCountDisjointIntervals(t *testing.T) {
	// define test case with an array of intervals and the expected result
	type testCase struct {
		num1     []int
		num2     []int
		expected int
	}
	// define test cases
	testCases := []testCase{
		{
			num1:     []int{2, 4, 3},
			num2:     []int{2, 4, 3},
			expected: 684,
		},
		{
			num1:     []int{8, 8, 8},
			num2:     []int{9, 9, 9, 9},
			expected: 10887,
		},
	}

	// loop through test cases and call the countDisjointIntervals function
	for idx, tc := range testCases {
		list1 := transformToListNode(tc.num1)
		list2 := transformToListNode(tc.num2)
		result := addTwoNumbers(list1, list2)
		// check the result against the expected result
		if toNumber(result) != tc.expected {
			t.Errorf("Error on test index [%v] Expected %d, got %d", idx, tc.expected, result)
		}
	}
}

func transformToListNode(value []int) *ListNode {
	var head *ListNode
	var previous *ListNode
	for i := 0; i < len(value); i++ {
		current := ListNode{Val: value[i]}
		if previous != nil {
			current.Next = previous
		}
		previous = &current
		head = &current
	}
	return head
}

func toNumber(node *ListNode) int {
	result := 0
	for node != nil {
		result += node.Val
		node = node.Next
	}
	return result
}
