package disjointIntervals

func countDisjointIntervals(intervals [][]int) int {
	count := 1
	previousInterval := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > previousInterval[1] {
			count++
		}
		if intervals[i][1] < previousInterval[1] {
			previousInterval = intervals[i]
		}
	}
	return count
}
