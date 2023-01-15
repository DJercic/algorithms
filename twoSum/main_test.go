package twoSum

import "testing"

func TestTwoSum(t *testing.T) {
	// define test case with an array of numbers, a target and the expected result
	type testCase struct {
		numbers  []int
		target   int
		expected []int
	}
	// define test cases
	testCases := []testCase{
		{numbers: []int{2, 7, 11, 15}, target: 13, expected: []int{0, 2}},
		{numbers: []int{3, 2, 4}, target: 5, expected: []int{0, 1}},
		{numbers: []int{2, 4, 8, 7, 9, 11}, target: 17, expected: []int{2, 4}},
		{numbers: []int{3, 3}, target: 6, expected: []int{0, 1}},
	}
	// loop through test cases
	for idx, tc := range testCases {
		result := twoSum(tc.numbers, tc.target)
		// check the result against the expected result
		if len(result) != len(tc.expected) {
			t.Errorf("On test [%v] expected %v, got %v", idx, tc.expected, result)
		}
		for i, v := range result {
			if v != tc.expected[i] {
				t.Errorf("On test [%v] expected %v, got %v", idx, tc.expected, result)
			}
		}
	}
}

/* {2, 4, 7, 8, 9, 10, 11 } target is 16

2+11 < 16
4+11 < 16
7+11 > 16
8+9 > 16
8+7 == 16
*/
