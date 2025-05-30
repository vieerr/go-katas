package main

import (
	"fmt"
	"math"
	"reflect"
	"sort"
)

func main() {

	a1 := []int{121, 144, 19, 161, 19, 144, 19, 11}
	a2 := []int{11 * 11, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19}
	fmt.Println(Comp(a1, a2))
	a1 = []int{121, 144, 19, 161, 19, 144, 19, 11}
	a2 = []int{121, 14641, 20736, 361, 25921, 361, 20736, 361}
	fmt.Println(Comp(a1, a2))
	a1 = []int{121, 144, 19, 161, 19, 144, 19, 11}
	a2 = []int{231, 14641, 20736, 361, 25921, 361, 20736, 361}
	fmt.Println(Comp(a1, a2))

	a1 = []int{121, 144, 19, 161, 19, 144, 19, 11}
	a2 = []int{121, 14641, 20736, 36100, 25921, 361, 20736, 361}
	fmt.Println(Comp(a1, a2))

	a1 = []int{}
	a2 = []int{}
	fmt.Println(Comp(a1, a2))

	a1 = []int{}
	a2 = []int{25, 49}
	fmt.Println(Comp(a1, a2))

	a1 = []int{121, 144, 19, 161, 19, 144, 19, 11, 1008}
	a2 = []int{121, 14641, 20736, 36100, 25921, 361, 20736, 361}
	fmt.Println(Comp(a1, a2))

	a1 = []int{-10000000, 100000000}
	a2 = []int{100000000000000, 10000000000000000}
	fmt.Println(Comp(a1, a2))

}

func Comp(array1 []int, array2 []int) bool {
	if array2 == nil || array1 == nil {
		return false
	}
	if len(array1) == 0 || len(array2) == 0 {
		return false
	}
	var sqrtB = make([]int, len(array2))
	for i, element := range array2 {
		sqrtEl := math.Sqrt(float64(element))
		sqrtB[i] = int(sqrtEl)
	}

	aMap := make(map[int]int)
	for _, el := range array1 {
		aMap[int(math.Abs(float64(el)))]++
	}
	bMap := make(map[int]int)
	for _, el := range sqrtB {
		bMap[el]++
	}
	fmt.Println(aMap)
	fmt.Println(bMap)

	eq := reflect.DeepEqual(aMap, bMap)
	if eq {
		return true
	} else {

		return false
	}
}

func CompACCEPTED(a, b []int) bool {
	if a == nil || b == nil {
		return false
	}

	c, d := a[:], b[:]
	for i, n := range a {
		c[i] = n * n
	}

	sort.Ints(c)
	sort.Ints(d)
	return reflect.DeepEqual(c, d)
}
