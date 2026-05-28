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

func TransposeMatrix(input [][]int) [][]int {
	rowLength := len(input)
	colLength := len(input[0])
	output := make([][]int, 0)
	//[row][column]
	//[  c1 c2 c3
	//   [1,2,3]  row 1
	//   [4,5,6]  row 2
	//   [7,8, 9]  row 3
	//]

	for i := 0; i < colLength; i++ {
		newRow := make([]int, 0)
		for j := 0; j < rowLength; j++ {
			//0  0,0 -> 1 , 1,0-> 4, 3,0-> 7
			//1  0,1 -> 2 , 1,1-> 5  2,1 -> 8
			newRow = append(newRow, input[j][i])
		}
		output = append(output, newRow)
	}

	return output
}

func ThreeNumberSum(array []int, target int) [][]int {
	// Write your code here.

	sort.Ints(array)
	response := make([][]int, 0)
	for key, cN := range array {

		left := key + 1
		right := len(array) - 1

		for left < right {
			sum := cN + array[left] + array[right]

			if sum == target {
				response = append(response, []int{cN, array[left], array[right]})
				left++
				right--
			} else if sum < target {
				left++
			} else {
				right--
			}
		}

	}
	return response
}

func SmallestDifference(array1, array2 []int) []int {
	sort.Ints(array1)
	sort.Ints(array2)

	pOne, pTwo := 0, 0
	smallestDiff := math.MaxInt
	bestPair := []int{}

	for pOne < len(array1) && pTwo < len(array2) {
		first := array1[pOne]
		second := array2[pTwo]

		diff := first - second
		if diff < 0 {
			diff = -diff
		}

		if diff < smallestDiff {
			smallestDiff = diff
			bestPair = []int{first, second}
		}

		if first < second {
			pOne++
		} else if second < first {
			pTwo++
		} else {
			return []int{first, second}
		}
	}

	return bestPair
}
