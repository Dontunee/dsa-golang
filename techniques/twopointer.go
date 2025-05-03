package techniques

func twoSum(inputArray []int, targetSum int) []int {
	pointerOne := 0
	pointerTwo := len(inputArray) - 1

	for pointerOne < pointerTwo {
		currentSum := inputArray[pointerOne] + inputArray[pointerTwo]

		if currentSum > targetSum {
			pointerTwo--
		} else if currentSum < targetSum {
			pointerOne++
		} else {
			return []int{pointerOne, pointerTwo}
		}
	}

	return nil
}
