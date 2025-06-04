package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(Permutations("aabb"))
}

func Contains(s []int, target int) bool {
	for _, v := range s {
		if v == target {
			return true
		}
	}
	return false
}

func Permutate(s []string, locked []int, combs map[string]struct{}) []string {
	if len(locked) == len(s) {
		pm := make([]string, len(s))
		for i, idx := range locked {
			pm[i] = s[idx]
		}
		k := strings.Join(pm, "")
		combs[k] = struct{}{}
	}
	for i := range s {
		if !Contains(locked, i) {
			// if !slices.Contains(locked, i) { // not supported in golang 1.20
			// newLocked := slices.Clone(locked) // not supported in golang 1.20
			newLocked := append([]int{}, locked...)
			newLocked = append(newLocked, i)
			Permutate(s, newLocked, combs)
		}

	}
	return nil
}

func Permutations(s string) (res []string) {
	combs := make(map[string]struct{})
	Permutate(strings.Split(s, ""), []int{}, combs)
	for k, _ := range combs {
		res = append(res, k)
	}
	return
}

func PermutationsACCEPTED(s string) (a []string) {
	if len(s) < 2 {
		return []string{s}
	}
	m := map[string]bool{}
	for _, sub := range Permutations(s[1:]) {
		for i := 0; i <= len(sub); i++ {
			st := sub[0:i] + s[0:1] + sub[i:]
			if m[st] {
				continue
			}
			m[st] = true
			a = append(a, st)
		}
	}
	return
}
