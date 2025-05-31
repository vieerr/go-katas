// 6 kyu

package main

import (
	"fmt"
	"sort"
	"strings"
)

// Description:

// John has invited some friends. His list is:

// s = "Fred:Corwill;Wilfred:Corwill;Barney:Tornbull;Betty:Tornbull;Bjon:Tornbull;Raphael:Corwill;Alfred:Corwill";

// Could you make a program that

//     makes this string uppercase
//     gives it sorted in alphabetical order by last name.

// When the last names are the same, sort them by first name. Last name and first name of a guest come in the result between parentheses separated by a comma.

// So the result of function meeting(s) will be:

// "(CORWILL, ALFRED)(CORWILL, FRED)(CORWILL, RAPHAEL)(CORWILL, WILFRED)(TORNBULL, BARNEY)(TORNBULL, BETTY)(TORNBULL, BJON)"

// It can happen that in two distinct families with the same family name two people have the same first name too.

func main() {
	fmt.Println(Meeting("Fred:Corwill;Wilfred:Corwill;Barney:Tornbull;Betty:Tornbull;Bjon:Tornbull;Raphael:Corwill;Alfred:Corwill"))
}

type Person struct {
	name  string
	lName string
}

func Meeting(s string) string {

	separated := strings.Split(s, ";")
	people := make([]Person, len(separated))
	for i, p := range separated {
		names := strings.Split(p, ":")
		people[i] = Person{name: strings.ToUpper(names[0]), lName: strings.ToUpper(names[1])}
	}
	sort.SliceStable(people, func(i, j int) bool {
		if strings.Compare(people[i].lName, people[j].lName) != 0 {
			return strings.Compare(people[i].lName, people[j].lName) == -1
		}
		return strings.Compare(people[i].name, people[j].name) == -1

	})
	var res string
	for _, p := range people {
		res += fmt.Sprintf("(%s, %s)", p.lName, p.name)
	}
	return res
}

func MeetingACCEPTED(s string) string {
	sw := strings.Split(strings.ToUpper(s), ";")
	res := []string{}
	for _, pn := range sw {
		a := strings.Split(pn, ":")
		s := "(" + a[1] + ", " + a[0] + ")"
		res = append(res, s)
	}
	sort.Strings(res)
	return strings.Join(res, "")
}
