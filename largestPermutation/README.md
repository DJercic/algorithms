# Largest permutation

Given an array A of a random permutation of numbers from 1 to N. Given B, the number of swaps in A that we can make.
  
Find the largest permutation possible

Constraints:
    - 1 <= N <= 10^6
    - 1 <= B <= 10^9

Examples

1. 
- A = [1,3,2], B=1 => expected: [3,1,2] (312 > 132 && 312 > 213)
- A = [1,2,3], B=1 => expected: [3,2,1]
- A = [2,2,3,6], B=3 => expected: [6,3,2,2]
- A = [3,2,4,1,5], B=3 => expected: [5,4,3,2,1]

Approach
Sort A in descending order. For each element in A, find the index of the largest element in A that is smaller than the current element. Swap the current element with the largest element. Decrement B by 1. If B is 0, return A.

[3,2,4,1,5]
[5,4,3,2,1]

  swaps=0, B=3
  1. [5,2,4,1,3] i=0, swaps=1, 
  2. [5,4,2,1,3] i=1, swaps=2,
  3. [5,4,3,1,2] i=3, swaps=3 