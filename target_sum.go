package main

// 1. Pair with Target Sum (in a sorted array)
// Problem:
// Given a sorted array of integers, find two numbers that sum up to a given target value.
// Return their indices or the numbers themselves.
// Example:
// Input: arr = [1, 2, 3, 4, 6], target = 6
// Output: (1, 3) (since 2 + 4 = 6)

func targetSum(input []int, sumToFind int) (int, int) {
	for i, j := 0, 1; i < j && j < len(input); {
		sum := input[i] + input[j]
		if sum < sumToFind {
			j++
		} else if sum == sumToFind {
			return i, j
		} else {
			i++
			j = i + 1
		}
	}

	return -1, -1
}
