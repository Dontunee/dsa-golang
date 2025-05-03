package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"

	"github.com/dontunee/dsa-golang/techniques"
)

func main() {
	fmt.Println(techniques.AllConstruct("purple", []string{"purp", "p", "ur", "le", "purpl"}))
}

func plusOne(digits []int) []int {
	length := len(digits)
	for i := length - 1; i >= 0; i++ {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}

	result := make([]int, length+1)
	result[0] = 1
	return result
}
func SortedSquares(nums []int) []int {
	length := len(nums)
	firstPointer := 0
	secondPointer := length - 1
	output := make([]int, length)
	counter := length - 1

	for firstPointer <= secondPointer {
		firstSquare := nums[firstPointer] * nums[firstPointer]
		secondSquare := nums[secondPointer] * nums[secondPointer]

		if firstSquare > secondSquare {
			output[counter] = firstSquare
			firstPointer++
		} else {
			output[counter] = secondSquare
			secondPointer--
		}
		counter--
	}

	return output

}

func ThreeSum(nums []int) [][]int {
	result := make([][]int, 0)
	sort.Ints(nums)

	for index, value := range nums {
		if index > 0 && value == nums[index-1] {
			continue
		}

		leftPointer, rightPointer := index+1, len(nums)-1

		for leftPointer < rightPointer {
			threeSum := value + nums[leftPointer] + nums[rightPointer]

			if threeSum > 0 {
				rightPointer--
			} else if threeSum < 0 {
				leftPointer++
			} else {
				newArray := []int{value, nums[leftPointer], nums[rightPointer]}
				result = append(result, newArray)
				leftPointer++

				for nums[leftPointer] == nums[leftPointer-1] && leftPointer < rightPointer {
					leftPointer += 1
				}

			}

		}
	}

	return result
}
func PrintTriangle(n int) {
	for i := 1; i <= n; i++ {
		for value := 0; value < i; value++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func PrintDaysOfTheWeek() {
	days := map[string]string{
		"mon":  "monday",
		"tue":  "tuesday",
		"wed":  "wednesday",
		"thur": "thursday",
		"fri":  "friday",
		"sat":  "saturday",
		"sun":  "sunday",
	}

	for key, value := range days {
		fmt.Println(key + " " + " stands for " + value)
	}
}

func MaximumGap(N int) int {
	maximumGap := math.MinInt
	currentCount := 0
	inputBinary := strconv.FormatInt(int64(N), 2)
	runeSample := []rune(inputBinary)
	for i := 0; i < len(runeSample); i++ {
		if runeSample[i] == '1' {
			i++
			for ; i < len(runeSample); i++ {
				if inputBinary[i] == '0' {
					currentCount++
				} else if inputBinary[i] == '1' {
					maximumGap = int(math.Max(float64(currentCount), float64(maximumGap)))
					currentCount = 0
				}
			}
		}
	}
	return maximumGap
}

func ReverseArray(input []int) []int {
	length := len(input)
	for i := 0; i < length/2; i++ {
		k := length - i - 1
		input[i], input[k] = input[k], input[i]
	}
	return input
}

func FindUnPairedElement(A []int) int {
	var arrayValue int
	dictionary := make(map[int]int, 0)
	for _, value := range A {
		dictionary[value]++
	}

	for _, arrayValue = range A {
		if figure, ok := dictionary[arrayValue]; ok {
			if figure == 1 {
				return arrayValue
			}
		}
	}

	return arrayValue
}

func ContainsDuplicate(nums []int) bool {
	length := len(nums)
	response := false
	if length < 1 || length > int(math.Pow(float64(10), float64(5))) {
		return false
	}

	dictionary := make(map[int]int, 0)
	for _, value := range nums {
		dictionaryValue := dictionary[value]
		if dictionaryValue > 0 {
			response = true
			return response
		}
		dictionary[value]++
	}
	return response
}

func RotateNTimes(A []int, K int) []int {
	// write your code in Go 1.4
	length := len(A)
	K = K % length
	reverse(A, 0, length-1)
	reverse(A, 0, K-1)
	reverse(A, K, length-1)
	return A
}

func reverse(input []int, start int, end int) []int {
	for start < end {
		temp := input[start]
		input[start] = input[end]
		input[end] = temp
		start++
		end--
	}
	return input
}

/*
create a responseArray and insert the first item of the input array into it
loop through from the second item to the last item in the input array
for each one of the items , pick it , then loop through the response array
compare it from the rightmost value
if the value in your hand is higher than the value i am comparing against skip it and move to the left
if it is higher , add it to the right of that value and move to the next value in the input array
if it is not higher than any , insert it at the beginning of the response array
then move to the next item in the input array
{4, 3, 2, 1}
{1,2,3,4}
*/
func insertionSort(inputArray []int) []int {
	if len(inputArray) < 2 {
		return inputArray
	}

	resp := make([]int, 1)
	resp[0] = inputArray[0]

	for i := 1; i < len(inputArray); i++ {
		numberAtHand := inputArray[i]
		j := len(resp) - 1

		// Find the right spot for numberAtHand
		for ; j >= 0 && resp[j] > numberAtHand; j-- {
			// No action needed here; just iterating to find the spot
		}

		// Insert numberAtHand after the found spot
		resp = append(resp, 0)       // Make room for the new element
		copy(resp[j+2:], resp[j+1:]) // Shift elements to the right
		resp[j+1] = numberAtHand     // Insert numberAtHand at the right spot
	}

	return resp
}

/*
{1,2,5,9,10} 22
{1,2,9,10}
*/

// Function to find a subset that sums up to the target
func findSubsetSum(S []int, T int) ([]int, bool) {
	n := len(S)
	// DP table to check if sum j can be achieved with subset of first i items
	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, T+1)
	}
	dp[0][0] = true // Base case: sum 0 is always achievable

	// Parent table to track the decisions
	parent := make([][]int, n+1)
	for i := range parent {
		parent[i] = make([]int, T+1)
		for j := range parent[i] {
			parent[i][j] = -1 // Initialize parent table with -1
		}
	}

	// Fill DP table and track decisions
	for i := 1; i <= n; i++ {
		for j := 0; j <= T; j++ {
			if j >= S[i-1] && dp[i-1][j-S[i-1]] {
				dp[i][j] = true
				parent[i][j] = S[i-1] // Record the item contributing to this sum
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	// If there's no subset that sums up to T, return an empty slice and false
	if !dp[n][T] {
		return nil, false
	}

	// Reconstruct the subset
	var subset []int
	for i, j := n, T; j > 0 && i > 0; i-- {
		if parent[i][j] != -1 {
			subset = append([]int{parent[i][j]}, subset...) // Prepend to maintain order
			j -= parent[i][j]
		}
	}

	return subset, true
}
