package highestProduct

import "testing"

// generate tests for the highestProduct function
func TestHighestProduct(t *testing.T) {
	// define test case with an array of numbers and the expected result
	type testCase struct {
		numbers  []int
		expected int
	}
	// define test cases
	testCases := []testCase{
		{numbers: []int{5, 1, 4, 3, 2}, expected: 60},
		{numbers: []int{1, 2, 3, 4, -5}, expected: 24},
		{numbers: []int{1, 2, -3, 4, -5}, expected: 60},
		{numbers: []int{1, 2, -3, -4, -5}, expected: 40},
		{numbers: []int{1, -2, -3, -4, -5}, expected: 20},
		{numbers: []int{-1, -2, -3, -4, -5}, expected: -6},
		//{numbers: []int{1}, expected: 1},
		//{numbers: []int{1, 2}, expected: 2},
		{numbers: []int{1, 2, 3}, expected: 6},
		{numbers: []int{1, 2, 3, 4}, expected: 24},
		{numbers: []int{1, 2, 3, 4, 5}, expected: 60},
		{numbers: []int{1, 2, 3, 4, 5, 6}, expected: 120},
		{numbers: []int{1, 2, 3, 4, 5, 6, 7}, expected: 210},
		{numbers: []int{1, 2, 3, 4, 5, 6, 7, 8}, expected: 336},
		{numbers: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, expected: 504},
		{numbers: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, expected: 720},
	}
	// loop through test cases
	for idx, tc := range testCases {
		result := highestProduct(tc.numbers)
		// check the result against the expected result
		if result != tc.expected {
			t.Errorf("On test [%v] expected %d, got %d", idx, tc.expected, result)
		}
	}

}
