package techniques

import (
	"fmt"
	"strings"
)

func fib(n int) int {
	if n <= 2 {
		return 1
	}

	return fib(n-1) + fib(n-2)
}

func FibMemo(n int) int {
	return fibMemoization(n, make(map[int]int, n))
}

func fibMemoization(n int, memo map[int]int) int {
	if memo[n] != 0 {
		return memo[n]
	}

	if n <= 2 {
		return 1
	}
	memo[n] = fibMemoization(n-1, memo) + fibMemoization(n-2, memo)
	return memo[n]
}

func GridTraveller(m int, n int) int {
	var memo = make(map[string]int)
	return gridTraveller(m, n, memo)
}

func gridTraveller(m int, n int, memo map[string]int) int {
	var key = fmt.Sprintf("%d,%d", m, n)

	if memo[key] != 0 {
		return memo[key]
	}
	if m == 0 || n == 0 {
		return 0
	}

	if m == 1 && n == 1 {
		return 1
	}

	memo[key] = gridTraveller(m-1, n, memo) + gridTraveller(m, n-1, memo)
	return memo[key]
}

func CanSum(targetSum int, numbers []int) bool {
	return canSum(targetSum, numbers, make(map[int]bool))
}

func canSum(targetSum int, numbers []int, memo map[int]bool) bool {
	value, ok := memo[targetSum]
	if ok {
		return value
	}

	if targetSum == 0 {
		return true
	}

	if targetSum < 0 {
		return false
	}

	for _, value := range numbers {
		var remainder = targetSum - value
		if canSum(remainder, numbers, memo) == true {
			memo[targetSum] = true
			return memo[targetSum]
		}
	}
	memo[targetSum] = false
	return memo[targetSum]
}

func HowSum(targetSum int, numbers []int) []int {
	if targetSum == 0 {
		return []int{}
	}

	if targetSum < 0 {
		return nil
	}

	for _, value := range numbers {
		var remainder = targetSum - value
		response := HowSum(remainder, numbers)
		if response != nil {
			response = append(response, value)
			return response
		}
	}

	return nil
}

func CanConstruct(target string, words []string) bool {
	if target == "" {
		return true
	}

	for _, word := range words {
		if strings.HasPrefix(target, word) {
			remainder := target[len(word):]
			if CanConstruct(remainder, words) == true {
				return true
			}
		}
	}
	return false
}

func CountConstruct(target string, words []string) int {
	if target == "" {
		return 1
	}

	count := 0

	for _, word := range words {
		if strings.HasPrefix(target, word) {
			var remainder = target[len(word):]
			if CountConstruct(remainder, words) == 1 {
				count++
			}
		}
	}

	return count
}

func AllConstruct(target string, wordBank []string) [][]string {
	if target == "" {
		return [][]string{{}} // Base case: return a slice containing an empty slice
	}

	var result [][]string

	for _, word := range wordBank {
		if strings.HasPrefix(target, word) {
			suffix := target[len(word):]                 // Get the remainder of the target after removing the prefix word
			suffixWays := AllConstruct(suffix, wordBank) // Recursively call AllConstruct on the suffix

			// Iterate over the ways to construct the suffix
			for _, way := range suffixWays {
				// Prepend the current word to each way
				prependWordToWay := append([]string{word}, way...)
				result = append(result, prependWordToWay) // Directly append to result
			}
		}
	}

	return result
}
