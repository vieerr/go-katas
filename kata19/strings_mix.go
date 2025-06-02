// 4 kyu

package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

// Description:

// Given two strings s1 and s2, we want to visualize how different the two strings are. We will only take into account the lowercase letters (a to z). First let us count the frequency of each lowercase letters in s1 and s2.

// s1 = "A aaaa bb c"

// s2 = "& aaa bbb c d"

// s1 has 4 'a', 2 'b', 1 'c'

// s2 has 3 'a', 3 'b', 1 'c', 1 'd'

// So the maximum for 'a' in s1 and s2 is 4 from s1; the maximum for 'b' is 3 from s2. In the following we will not consider letters when the maximum of their occurrences is less than or equal to 1.

// We can resume the differences between s1 and s2 in the following string: "1:aaaa/2:bbb" where 1 in 1:aaaa stands for string s1 and aaaa because the maximum for a is 4. In the same manner 2:bbb stands for string s2 and bbb because the maximum for b is 3.

// The task is to produce a string in which each lowercase letters of s1 or s2 appears as many times as its maximum if this maximum is strictly greater than 1; these letters will be prefixed by the number of the string where they appear with their maximum value and :. If the maximum is in s1 as well as in s2 the prefix is =:.

// In the result, substrings (a substring is for example 2:nnnnn or 1:hhh; it contains the prefix) will be in decreasing order of their length and when they have the same length sorted in ascending lexicographic order (letters and digits - more precisely sorted by codepoint); the different groups will be separated by '/'. See examples and "Example Tests".

// Hopefully other examples can make this clearer.

// s1 = "my&friend&Paul has heavy hats! &"
// s2 = "my friend John has many many friends &"
// mix(s1, s2) --> "2:nnnnn/1:aaaa/1:hhh/2:mmm/2:yyy/2:dd/2:ff/2:ii/2:rr/=:ee/=:ss"

// s1 = "mmmmm m nnnnn y&friend&Paul has heavy hats! &"
// s2 = "my frie n d Joh n has ma n y ma n y frie n ds n&"
// mix(s1, s2) --> "1:mmmmmm/=:nnnnnn/1:aaaa/1:hhh/2:yyy/2:dd/2:ff/2:ii/2:rr/=:ee/=:ss"

// s1="Are the kids at home? aaaaa fffff"
// s2="Yes they are here! aaaaa fffff"
// mix(s1, s2) --> "=:aaaaaa/2:eeeee/=:fffff/1:tt/2:rr/=:hh"

// Note for Swift, R, PowerShell

// The prefix =: is replaced by E:

// s1 = "mmmmm m nnnnn y&friend&Paul has heavy hats! &"
// s2 = "my frie n d Joh n has ma n y ma n y frie n ds n&"
// mix(s1, s2) --> "1:mmmmmm/E:nnnnnn/1:aaaa/1:hhh/2:yyy/2:dd/2:ff/2:ii/2:rr/E:ee/E:ss"

func main() {
	// s1 := "mmmmm m nnnnn y&friend&Paul has heavy hats! &"
	// s2 := "my frie n d Joh n has ma n y ma n y frie n ds n&"
	// Mix(s1, s2)
	r := Mix("looping is fun but dangerous", "less dangerous than coding")
	fmt.Println(r)
}

func Mix(s1, s2 string) (res string) {

	mapS1 := make(map[string]int)
	mapS2 := make(map[string]int)

	for _, c := range s1 {
		if unicode.IsLetter(c) && unicode.IsLower(c) {
			mapS1[string(c)]++

		}
	}
	for _, c := range s2 {
		if unicode.IsLetter(c) && unicode.IsLower(c) {
			mapS2[string(c)]++
		}
	}

	var totals []string

	for k, v := range mapS1 {
		if v > 1 {
			if mapS1[k] > mapS2[k] {
				mapS2[k] = 0
				totals = append(totals, "1:"+strings.Repeat(k, v))
			}
			if mapS1[k] == mapS2[k] {
				mapS2[k] = 0
				totals = append(totals, "=:"+strings.Repeat(k, v))
			}
		}
	}
	for k, v := range mapS2 {
		if v > 1 {
			totals = append(totals, "2:"+strings.Repeat(k, v))
		}
	}

	sort.Strings(totals)
	sort.SliceStable(totals, func(i, j int) bool {
		if len(totals[i]) == len(totals[j]) {
			return totals[i] < totals[j]
		}
		return len(totals[i]) > len(totals[j])
	})

	res = strings.Join(totals, "/")
	return
}

func MixACCEPTED(s1, s2 string) string {
	alphabase := "abcdefghijklmnopqrstuvwxyz"
	result := []string{}
	for _, c := range alphabase {
		nb_s1 := strings.Count(s1, string(c))
		nb_s2 := strings.Count(s2, string(c))
		if nb_s1 > 1 || nb_s2 > 1 {
			if nb_s1 == nb_s2 {
				result = append(result, "=:"+strings.Repeat(string(c), nb_s1))
			}
			if nb_s1 > nb_s2 {
				result = append(result, "1:"+strings.Repeat(string(c), nb_s1))
			}
			if nb_s1 < nb_s2 {
				result = append(result, "2:"+strings.Repeat(string(c), nb_s2))
			}
		}
	}
	sort.Slice(result, func(i, j int) bool {
		if len(result[i]) == len(result[j]) {
			return result[i] < result[j]
		}
		return len(result[i]) > len(result[j])
	})
	return strings.Join(result, "/")
}
