# Disjoint intervals

### Level: Beginner

### Problem Statement

Given a list of intervals: `[start, end]`
find the length of maximal set of mutually disjoint intervals.

- Constraints:
    - 1 <= N <= 1e5
    - A[i] = [1, 1e9]

### Examples:
 
 - [[1,2], [2,10], [4,6]] => output: 2
   - [1,2] and [4,6] are disjoint
 - [[2,5], [3,7], [6,8]] => output: 2
   - [2,5] and [6,8] are disjoint
 - [[1,2], [2,3], [3,4]] => output: 2
   - [1,2] and [3,4] are disjoint
 - [[1,2], [2,3], [3,4], [4,5]] => output: 0
   - no intervals are disjoint
- [[1,5], [2,3], [4,6], [7,9]] => output: 3
    - [2,3], [4,6], [7,9] are disjoint

### Solution:
    - Input: [[3,10], [2,5], [8, 12], [4,7]]
    - Input: [[1,4], [2,3], [4,6], [8,9]]

sort by start (n * logn) => [[2,5], [3,4], [5,7], [8,12]]
if the first interval's end time is less than the second interval's start time, then they are disjoint

```python
count = 1   # first interval is always disjoint
previousInterval = intervals[0]
for i in range(1, len(intervals)):
    if previuosInterval[1] < intervals[i][0]:
        count += 1
        previousInterval = intervals[i]
    if previuosInterval[1] > intervals[i][1]:
        previousInterval = intervals[i]
```
