// 4 kyu
package main

import (
	"fmt"
	"math"
	"strconv"
)

// Lyrics...

// Pyramids are amazing! Both in architectural and mathematical sense. If you have a computer, you can mess with pyramids even if you are not in Egypt at the time. For example, let's consider the following problem. Imagine that you have a pyramid built of numbers, like this one here:

//    /3/
//   \7\ 4
//  2 \4\ 6
// 8 5 \9\ 3

// Here comes the task...

// Let's say that the 'slide down' is the maximum sum of consecutive numbers from the top to the bottom of the pyramid. As you can see, the longest 'slide down' is 3 + 7 + 4 + 9 = 23

// Your task is to write a function that takes a pyramid representation as an argument and returns its largest 'slide down'. For example:

// * With the input `[[3], [7, 4], [2, 4, 6], [8, 5, 9, 3]]`
// * Your function should return `23`.

// By the way...

// My tests include some extraordinarily high pyramids so as you can guess, brute-force method is a bad idea unless you have a few centuries to waste. You must come up with something more clever than that.

func main() {
	res := LongestSlideDown([][]int{{1}, {2, 3}, {1, 3, 30}, {50, 1, 1, 1}})
	// res := LongestSlideDown([][]int{{3}, {7, 4}, {2, 1, 10}, {28, 5, 9, 3}})
	fmt.Println(res)
}

type Walked struct {
	walkers map[string]int
}

func walker(pyramid [][]int, parentI, parentJ int, walked *Walked) (res int) {
	if parentI == len(pyramid)-1 {
		return pyramid[parentI][parentJ]
	}

	leftKey := strconv.Itoa(parentI+1) + strconv.Itoa(parentJ)
	rightKey := strconv.Itoa(parentI+1) + strconv.Itoa(parentJ+1)

	left, okL := walked.walkers[leftKey]
	right, okR := walked.walkers[rightKey]
	if okL && okR {
		return pyramid[parentI][parentJ] + int(math.Max(float64(left), float64(right)))
	}

	walked.walkers[leftKey] = walker(pyramid, parentI+1, parentJ, walked)
	walked.walkers[rightKey] = walker(pyramid, parentI+1, parentJ+1, walked)
	return pyramid[parentI][parentJ] + int(math.Max(float64(walked.walkers[leftKey]), float64(walked.walkers[rightKey])))
}

func LongestSlideDown(pyramid [][]int) (res int) {
	var walked Walked
	walked.walkers = make(map[string]int)
	res = walker(pyramid, 0, 0, &walked)
	return
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func LongestSlideDownACCEPTED(pyramid [][]int) int {
	for i := len(pyramid) - 2; i >= 0; i-- {
		for j := 0; j < len(pyramid[i]); j++ {
			pyramid[i][j] += max(pyramid[i+1][j], pyramid[i+1][j+1])
		}
	}
	return pyramid[0][0]
}
