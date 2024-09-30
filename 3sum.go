package main

import (
	"slices"
)

func sumThree(input []int, targetSum int) [][3]int {
	slices.Sort(input)
	n := len(input)
	cnt := 0
	var res [][3]int

	i := 0

	for ; i < n-2 && input[i] < 0; i++ {
		if i > 0 && input[i] == input[i-1] {
			continue
		}

	left_right:
		for l, r := i+1, n-1; l < r && r < n; {
			if l > i+1 && input[l] == input[l-1] {
				l++
				continue
			}

			cnt++
			sum := input[l] + input[r] + input[i]

			for sum < targetSum && (input[l+n/10]+input[r]+input[i] < targetSum) {
				l += n / 5
				continue left_right
			}

			if sum < targetSum {
				l++
			} else if sum == targetSum {
				res = append(res, [3]int{input[i], input[l], input[r]})
				r--
				l++
			} else {
				r--
			}
		}

	}

	return res
}

func sumThreeSlow(input []int, target int) [][3]int {
	cnt := 0
	n := len(input)
	slices.Sort(input)
	var res [][3]int
	for i := 0; i < n-2; i++ {
	left:
		for l := i + 1; l < n-1; l++ {
			if input[l] == input[l-1] {
				continue
			}
			for r := n - 1; r > l; r-- {
				if input[r] == input[r-1] {
					continue
				}

				sum := input[i] + input[l] + input[r]
				cnt++
				if sum < target {
					continue left
				}

				if sum == target {
					res = append(res, [3]int{input[i], input[l], input[r]})
				} else if sum > target {
					continue
				}
			}

		}
	}

	return res
}
