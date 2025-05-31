// 5 kyu

package main

import (
	"fmt"
	"sort"
	"strings"
)

// My friend John and I are members of the "Fat to Fit Club (FFC)". John is worried because each month a list with the weights of members is published and each month he is the last on the list which means he is the heaviest.

// I am the one who establishes the list so I told him: "Don't worry any more, I will modify the order of the list". It was decided to attribute a "weight" to numbers. The weight of a number will be from now on the sum of its digits.

// For example 99 will have "weight" 18, 100 will have "weight" 1 so in the list 100 will come before 99.

// Given a string with the weights of FFC members in normal order can you give this string ordered by "weights" of these numbers?

// Example:
// "56 65 74 100 99 68 86 180 90" ordered by numbers weights becomes:

// "100 180 90 56 65 74 68 86 99"
// When two numbers have the same "weight", let us class them as if they were strings (alphabetical ordering) and not numbers:

// 180 is before 90 since, having the same "weight" (9), it comes before as a string.

// All numbers in the list are positive numbers and the list can be empty.

// Notes
// it may happen that the input string have leading, trailing whitespaces and more than a unique whitespace between two consecutive numbers

func main() {
	fmt.Println(OrderWeight("2000    10003 1234000 44444444 9999 11 11 22 123"))
}

func OrderWeight(strng string) string {
	wArr := strings.Fields(strng)

	calcs := make([]int, len(wArr))
	weightMap := make(map[string]int)
	for i, weight := range wArr {
		_, ok := weightMap[string(weight)]
		if !ok {

			for _, char := range weight {
				calcs[i] += int(char - '0')
				weightMap[string(weight)] += int(char - '0')

			}
		}
	}
	sort.Slice(wArr, func(i, j int) bool {
		if weightMap[wArr[i]] < weightMap[wArr[j]] {
			return true
		}
		if weightMap[wArr[i]] > weightMap[wArr[j]] {
			return false
		}
		if weightMap[wArr[i]] == weightMap[wArr[j]] {
			return strings.Compare(wArr[i], wArr[j]) == -1
		}
		return strings.Compare(wArr[i], wArr[j]) == -1
	})

	return strings.Join(wArr, " ")
}

func OrderWeightACCEPTED(s string) string {
	// convert to arr
	arr := strings.Fields(s)

	// sort
	sort.SliceStable(arr, func(i, j int) bool {
		if diff := stringSum(arr[i]) - stringSum(arr[j]); diff == 0 { // if same "weight"
			return arr[i] < arr[j] // just compare the strings directly
		} else {
			return diff < 0 // otherwise compare using weight diff
		}
	})

	// convert back to string
	return strings.Join(arr, " ")
}
func stringSum(s string) int { // calculates sum of str's digits
	sum := 0
	for _, v := range s {
		sum += int(v) - '0'
	}
	return sum
}
