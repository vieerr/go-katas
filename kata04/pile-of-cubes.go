package main

import "math"

func main() {
	println(FindNb(1071225))
}

func FindNb(m int) int {
	n := (-1 + math.Sqrt(1+(8*math.Sqrt(float64(m))))) / 2
	if n == math.Trunc(n) {
		return int(n)
	} else {
		return -1
	}
}

func FindNbACCEPTED(m int) int {
	for n := 1; m > 0; n++ {
		m -= n * n * n
		if m == 0 {
			return n
		}
	}
	return -1
}
