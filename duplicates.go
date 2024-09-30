package main

// 2. Remove Duplicates from Sorted Array
// Problem:
// Given a sorted array, remove all rmDup in place such that each element appears only once
// and return the length of the modified array.
// Example:
// Input: arr = [1, 1, 2, 2, 3, 4, 4]
// Output: 4 (array becomes [1, 2, 3, 4])
func rmDup(input []int) []int {
	var result []int
	var lastAdded int
	// ops := 0

	for i, val := range input {
		if val != lastAdded {
			result = append(result, input[i])
		}
		lastAdded = input[i]
	}

	// for i, j := 0, 1; i < j && j < len(input); j++ {
	// 	ops++
	// 	if input[i] != lastAdded {
	// 		lastAdded = input[i]
	// 		result = append(result, lastAdded)
	// 		if input[i] == input[j] {
	// 			i += 2
	// 			j++
	// 		} else {
	// 			i = j
	// 		}
	// 	} else {
	// 		i = j
	// 	}
	// }

	return result
}
