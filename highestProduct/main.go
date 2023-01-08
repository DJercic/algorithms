package highestProduct

import (
	"sort"
)

func highestProduct(numbers []int) int {
	// M is multiplication factor.
	// Example: if M = 3 then we are looking for the highest product of 3 numbers.
	// M := 3
	sort.Sort(sort.IntSlice(numbers))
	withNegative := numbers[0] * numbers[1] * numbers[len(numbers)-1]
	withPositive := numbers[len(numbers)-1] * numbers[len(numbers)-2] * numbers[len(numbers)-3]
	if withNegative > withPositive {
		return withNegative
	} else {
		return withPositive
	}
}
