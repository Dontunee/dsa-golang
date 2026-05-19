package techniques

import (
	"math"
	"reflect"
)

func FindAnagrams(s string, p string) []int {
	result := make([]int, 0)
	pArray := []rune(p)
	sArray := []rune(s)
	pMap := make(map[rune]int, 0)
	customMap := make(map[rune]int, 0)
	left := 0
	right := 0
	pCount := 0
	length := len(pArray)

	for _, val := range pArray {
		pMap[val]++
	}

	for right <= len(sArray) {
		if pCount == length {
			if reflect.DeepEqual(pMap, customMap) {
				result = append(result, left)
			}
			if customMap[sArray[left]] > 1 {
				customMap[sArray[left]] -= 1
			} else if customMap[sArray[left]] == 1 {
				delete(customMap, sArray[left])
			}
			left += 1
			pCount -= 1
		}

		if right < len(sArray) {
			customMap[sArray[right]]++
		}
		right++
		pCount++
	}
	return result
}

func MaxSubArray(nums []int) int {
	maximumSum := math.MinInt64
	currentSum := 0
	windowStart := 0

	for windowEnd := 0; windowEnd < len(nums); windowEnd++ {
		currentSum += nums[windowEnd]

		if currentSum < 0 {
			maximumSum = int(math.Max(float64(currentSum), float64(maximumSum)))
			currentSum -= nums[windowStart]
			windowStart++
		} else if nums[windowEnd] > currentSum {
			currentSum = nums[windowEnd]
		} else {
			maximumSum = int(math.Max(float64(currentSum), float64(maximumSum)))
		}

	}

	return maximumSum
}
func BruteForceAverageOfKSubArray(input []int, k int) []float32 {
	var solution []float32
	var length int = len(input)
	difference := length - k

	for i := 0; i <= difference; i++ {
		average := float32(input[i]+input[i+1]+input[i+2]+input[i+3]+input[i+4]) / float32(k)
		solution = append(solution, average)
	}
	return solution
}

func AverageOfKSubArray(input []int, k int) []float32 {
	var result []float32
	currentSum := 0

	for i := 0; i < len(input); i++ {
		currentSum += input[i]

		if i >= k-1 {
			average := float32(currentSum) / float32(k)
			result = append(result, average)
			currentSum -= input[i-(k-1)]
		}
	}
	return result
}

func FindMaxSumSubArray(input []int, k int) int {
	maximumSum := math.MinInt
	currentSum := 0
	for i := 0; i < len(input); i++ {
		currentSum += input[i]

		if i >= k-1 {
			maximumSum = int(math.Max(float64(maximumSum), float64(currentSum)))
			currentSum -= input[i-(k-1)]
		}
	}
	return maximumSum
}

//Returns the size of the smallest subarray that is greater than or equal to target sum

func SmallestSubArraySize(inputArray []int, targetSum int) int {
	var smallestSize int = math.MaxInt
	var windowStart int
	var currentSum int

	for windowEnd := 0; windowEnd < len(inputArray); windowEnd++ {
		currentSum += inputArray[windowEnd]

		for currentSum >= targetSum {
			smallestSize = int(math.Min(float64(smallestSize), float64(windowEnd-windowStart+1)))
			currentSum -= inputArray[windowStart]
			windowStart++
		}
	}
	return smallestSize
}

func LongestSubstringKDistinct(input string, k int) int {
	var longest int = math.MinInt
	hashMap := make(map[rune]int)
	inputArray := []rune(input)
	var windowStart int
	var currentLength int

	for windowEnd := 0; windowEnd < len(inputArray); windowEnd++ {
		hashMap[inputArray[windowEnd]] += 1
		currentLength++
		for len(hashMap) > k {
			if hashMap[inputArray[windowStart]] == 1 {
				delete(hashMap, inputArray[windowStart])
			} else {
				hashMap[inputArray[windowStart]]--
			}
			currentLength--
			windowStart++
		}

		longest = int(math.Max(float64(longest), float64(currentLength)))
	}
	return longest
}
