// 5 kyu

package main

import "fmt"

// The Fibonacci numbers are the numbers in the following integer sequence (Fn): 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, ...

// such that:
// F(0)=0F(1)=1F(n)=F(n−1)+F(n−2)F(0) = 0\\F(1) = 1\\F(n) = F(n-1) + F(n-2)F(0)=0F(1)=1F(n)=F(n−1)+F(n−2)

// Given a number, say prod (for product), we search two Fibonacci numbers F(n) and F(n+1) verifying:
// F(n)∗F(n+1)=prodF(n) * F(n+1) = prodF(n)∗F(n+1)=prod

// Your function takes an integer (prod) and returns an array/tuple (check the function signature/sample tests for the return type in your language):

//     if F(n) * F(n+1) = prod:

//     (F(n), F(n+1), true)

//     If you do not find two consecutive F(n) verifying F(n) * F(n+1) = prod:

//     (F(n), F(n+1), false)

//     where F(n) is the smallest one such as F(n) * F(n+1) > prod.

// Examples:

// 714 ---> (21, 34, true)
// --> since F(8) = 21, F(9) = 34 and 714 = 21 * 34

// 800 --->  (34, 55, false)
// --> since F(8) = 21, F(9) = 34, F(10) = 55 and 21 * 34 < 800 < 34 * 55

func main() {
	res := ProductFib(84049690)
	fmt.Println(res)
}

func fib(n int, stored map[int]int) int {
	if n == 0 || n == 1 {
		return n
	}

	a, okA := stored[n-2]
	b, okB := stored[n-1]
	if !okA {
		stored[n-2] = fib(n-2, stored)
	}
	if !okB {
		stored[n-1] = fib(n-1, stored)
	}

	return a + b
}

func ProductFib(prod uint64) (res [3]uint64) {
	stored := make(map[int]int)
	i := 0
	for {
		a := uint64(fib(i, stored))
		b := uint64(fib(i+1, stored))
		calc := a * b
		if calc == prod {
			res[0] = b
			res[1] = a
			res[2] = 1
			break
		}
		if calc > prod {
			res[0] = b
			res[1] = a
			res[2] = 0
			break
		}
		i++

	}
	return res
}

func ProductFibACCEPTED(prod uint64) [3]uint64 {
	f1 := uint64(0)
	f2 := uint64(1)
	for f1*f2 < prod {
		f1, f2 = f2, f1+f2
	}
	success := uint64(0)
	if prod == f1*f2 {
		success = 1
	}
	return [3]uint64{f1, f2, success}
}
