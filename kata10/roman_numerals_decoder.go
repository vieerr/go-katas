// 6 kyu

package main

import (
	"fmt"
)

// Create a function that takes a Roman numeral as its argument and returns its value as a numeric decimal integer. You don't need to validate the form of the Roman numeral.

// Modern Roman numerals are written by expressing each decimal digit of the number to be encoded separately, starting with the leftmost digit and skipping any 0s. So 1990 is rendered "MCMXC" (1000 = M, 900 = CM, 90 = XC) and 2008 is rendered "MMVIII" (2000 = MM, 8 = VIII). The Roman numeral for 1666, "MDCLXVI", uses each letter in descending order.
// Example:

// "MM"      -> 2000
// "MDCLXVI" -> 1666
// "M"       -> 1000
// "CD"      ->  400
// "XC"      ->   90
// "XL"      ->   40
// "I"       ->    1

// Help:

// Symbol    Value
// I          1
// V          5
// X          10
// L          50
// C          100
// D          500
// M          1,000

func main() {
	fmt.Println(Decode("IV"))
}

func Decode(roman string) int {
	letters := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	res := 0
	for i, char := range roman {
		if i < (len(roman)-1) && letters[string(char)] < letters[string(roman[i+1])] {
			res -= letters[string(char)]

		} else {
			res += letters[string(char)]
		}

	}
	return res
}

func DecodeACCEPTED(roman string) int {
	var intern []int
	var result int
	for _, r := range roman {
		intern = append(intern, mapping[r])
	}

	for i := 1; i < len(intern); i++ {
		if intern[i-1] >= intern[i] {
			result += intern[i-1]
		} else {
			result -= intern[i-1]
		}
	}

	result += intern[len(intern)-1]

	return result
}

var mapping map[rune]int = map[rune]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
