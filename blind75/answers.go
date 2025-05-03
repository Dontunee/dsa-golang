package blind75

import "math"

func TwoSum(nums []int, target int) []int {
	//loop through array
	//if there is no number that the difference to target is there already
	//add it
	//else return the numbers index and current index
	length := len(nums)
	result := make([]int, 0)

	if length < 2 || length > int(math.Pow(10, 4)) {
		return result
	}

	if target < -int(math.Pow(10, 9)) || target > int(math.Pow(10, 9)) {
		return result
	}

	dictionary := make(map[int]int, 0)
	for index, value := range nums {
		differenceFrom := target - value
		if ind, ok := dictionary[differenceFrom]; ok {
			result = append(result, ind)
			result = append(result, index)
			return result
		}

		dictionary[value] = index
	}
	return result
}

func MaxProfit(prices []int) int {
	//have 2 pointers leftWindow and rightWindow
	//if rightWindow val is less then leftWindow val
	//shift leftWindow forward

	//else if it is greater // check the difference between the value
	//update the highest difference compared to current highest difference

	//return current highest difference
	length := len(prices)
	if length < 1 || length > int(math.Pow(10, 5)) {
		return 0
	}
	leftPointer, rightPointer := 0, 1
	currentMaximum := 0

	for rightPointer < length {
		leftValue := prices[leftPointer]
		rightValue := prices[rightPointer]

		if rightValue < leftValue {
			leftPointer++
		} else {
			difference := prices[rightPointer] - prices[leftPointer]
			currentMaximum = int(math.Max(float64(difference), float64(currentMaximum)))
			rightPointer++
		}
	}

	return currentMaximum

}

func ContainsDuplicate(nums []int) bool {
	//loop through
	//if any element exist before
	//return true
	//else
	//add to the dicti0nary

	length := len(nums)

	if length < 1 || length > int(math.Pow(10, 5)) {
		return false
	}

	dictionary := make(map[int]int, 0)
	for _, value := range nums {
		if _, ok := dictionary[value]; ok {
			return true
		}
		dictionary[value]++
	}

	return false
}