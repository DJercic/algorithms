# Lights on 

Given N bulbs, bulb can be turned on (1) or off (0). Turning on a bulb causes all remaining bulbs to the right to flip.

Find the min number of switches to run all the bulbs on.

- Constraints:
	- 1 <= N <= 1e5
	- A[i] = {0, 1}


Example:
- [1, 1, 1, 1, 1] => expected: 0
- [0, 0, 0, 0, 0] => expected: 1
- [1, 0, 1, 0, 1] => expected: 4
- [1, 0, 0, 0, 1] => expected: 2
