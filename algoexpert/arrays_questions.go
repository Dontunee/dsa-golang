package algoexpert

import "sort"

func NonConstructibleChange(coins []int) int {

	if len(coins) < 1 {
		return 1
	}
	sort.Ints(coins)
	cum := 0
	for i := 0; i < len(coins); i++ {
		if coins[i] > cum+1 {
			return cum + 1
		}
		cum += coins[i]
	}

	return cum + 1
}
