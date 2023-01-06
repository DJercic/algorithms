package lightsOn

import "testing"

// Generate tests for the turnOn function
func TestTurnOn(t *testing.T) {
	// Define test case with an array of bulbs and the expected result
	type testCase struct {
		bulbs    []int
		expected int
	}
	// Define test cases
	testCases := []testCase{
		{bulbs: []int{1, 1, 1, 1, 1}, expected: 0},
		{bulbs: []int{0, 0, 0, 0, 0}, expected: 1},
		{bulbs: []int{1, 0, 1, 0, 1}, expected: 4},
		{bulbs: []int{1, 0, 0, 0, 1}, expected: 2},
		{bulbs: []int{0, 1, 0, 1, 0}, expected: 5},
		{bulbs: []int{0, 1, 0, 1, 1, 1, 0, 1, 0, 1}, expected: 8},
	}
	// Loop through test cases
	for _, tc := range testCases {
		// Call the turnOn function
		result := turnOn(tc.bulbs)
		// Check the result against the expected result
		if result != tc.expected {
			t.Errorf("Expected %d, got %d", tc.expected, result)
		}
	}
}
