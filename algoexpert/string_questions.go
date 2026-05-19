package algoexpert

import (
	"sort"
	"strconv"
	"strings"
)

func IsPalindrome(str string) bool {
	// Write your code here.
	return buildFromBack(str)
}

func buildFromBack(str string) bool {
	strArray := []rune(str)
	if len(strArray) <= 0 {
		return false
	}
	var builder strings.Builder
	for i := len(strArray) - 1; i >= 0; i-- {
		builder.WriteString(string(strArray[i]))
	}

	if str == builder.String() {
		return true
	}

	return false
}

// O(n) time complexity and O(n) space complexity
func recursive(str string, i, j int) bool {
	//base case
	if i >= j {
		return true
	}

	//General case
	if str[i] != str[j] {
		return false
	}

	return recursive(str, i+1, j-1)
}

// O(n) time complexity and O(1) space complexity
func pointer(str string) bool {
	// Write your code here.
	j := len(str) - 1
	for i := 0; i < len(str); i++ {
		if str[i] != str[j] {
			return false
		}
		j--
	}

	return true
}

func CaesarCipherEncryptor(str string, key int) string {
	// Write your code here.
	solution := caesarCipherEncryptor(str, key)
	return solution
}

func caesarCipherEncryptor(str string, key int) string {
	const alphabets = "abcdefghijklmnopqrstuvwxyz"
	alphaArray := []rune(alphabets)

	sAR := []rune(str)
	var builder strings.Builder
	for _, v := range sAR {
		indexOf := indexOf(v, alphaArray)

		newIndex := (indexOf + key) % 26
		builder.WriteRune(alphaArray[newIndex])
	}
	return builder.String()
}

func indexOf(r rune, input []rune) int {
	for p, v := range input {
		if v == r {
			return p
		}
	}
	return 0
}

func RunLengthEncoding(str string) string {
	// Write your code here.
	return runLengthEncoding(str)
}

func runLengthEncoding(str string) string {

	counter := 1
	var sb strings.Builder
	for i := 1; i < len(str); i++ {
		currentChar := str[i]
		prevChar := str[i-1]

		if currentChar != prevChar || counter == 9 {
			//add count to characters
			//add previousChar to characters
			sb.WriteString(strconv.Itoa(counter))
			sb.WriteRune(rune(prevChar))
			counter = 0
		}
		counter++
	}

	//add count to characters
	//add last element to characters
	sb.WriteString(strconv.Itoa(counter))
	sb.WriteRune(rune(str[len(str)-1]))

	//conv characters to string
	return sb.String()
}

func FirstNonRepeatingCharacter(str string) int {
	// Write your code here.

	//abcdcaf
	d := map[rune]int{}

	for _, v := range str {
		d[v]++
	}

	for key, value := range str {
		if d[value] == 1 {
			return key
		}
	}
	return -1
}

func LongestPalindromicSubstring(str string) string {
	// Write your code here.

	currentLongest := []int{0, 1}

	for i := 1; i <= len(str); i++ {
		// get longest by checking the 2 sides in an odd string
		odd := getLongestPalindromeFrom(str, i-1, i+1)

		//get longest by checking 2 sdes in an even string which is basically putting cursor
		// on that space
		even := getLongestPalindromeFrom(str, i-1, i)

		// get the maximum by subtracting what is in array index 1 and index 0 of add and even
		// check the highest value retrieved and return the one with it
		longest := max(odd, even)
		currentLongest = max(longest, currentLongest)
	}
	return str[currentLongest[0]:currentLongest[1]]
}

func getLongestPalindromeFrom(str string, leftIdx, rightIdx int) []int {
	for leftIdx >= 0 && rightIdx < len(str) {
		// if the value at the left idx and the value at the right idx
		//are not the same , break out
		if str[leftIdx] != str[rightIdx] {
			break
		}
		leftIdx -= 1
		rightIdx += 1
	}
	return []int{leftIdx + 1, rightIdx}
}

func max(firstArray, secondArray []int) []int {

	fA1 := firstArray[1]
	fA0 := firstArray[0]
	differenceFirstArray := fA1 - fA0

	sA1 := secondArray[1]
	sA0 := secondArray[0]
	differenceSecondArray := sA1 - sA0

	if differenceFirstArray > differenceSecondArray {
		return firstArray
	}
	return secondArray
}

func GroupAnagrams(words []string) [][]string {
	// Write your code here.
	return groupAnagrams(words)
}

func groupAnagrams(words []string) [][]string {
	result := [][]string{}

	if len(words) < 1 {
		return result
	}
	d := map[string][]string{}

	arranged := rearrangeStrings(words)

	for k, v := range arranged {
		if d[v] != nil {
			d[v] = append(d[v], words[k])
			continue
		}
		d[v] = []string{words[k]}
	}

	for _, v := range d {
		result = append(result, v)
	}

	return result
}

func rearrangeStrings(str []string) []string {
	var result []string
	for _, v := range str {
		s := strings.Split(v, "")
		sort.Strings(s)
		join := strings.Join(s, "")
		result = append(result, join)
	}
	return result
}
