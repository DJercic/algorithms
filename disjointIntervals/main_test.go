package disjointIntervals

import "testing"

func TestCountDisjointIntervals(t *testing.T) {
	// define test case with an array of intervals and the expected result
	type testCase struct {
		intervals [][]int
		expected  int
	}
	// define test cases
	testCases := []testCase{
		{intervals: [][]int{{1, 2}, {3, 4}, {5, 6}}, expected: 3},
		{intervals: [][]int{{1, 4}, {2, 3}, {4, 5}}, expected: 2},
		{intervals: [][]int{{1, 5}, {2, 3}, {3, 5}}, expected: 1},
	}

	// loop through test cases and call the countDisjointIntervals function
	for idx, tc := range testCases {
		result := countDisjointIntervals(tc.intervals)
		// check the result against the expected result
		if result != tc.expected {
			t.Errorf("Error on test index [%v] Expected %d, got %d", idx, tc.expected, result)
		}
	}
}
