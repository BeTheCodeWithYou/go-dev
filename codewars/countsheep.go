package main

import (
	"strconv"
)

/*
*
Given a non-negative integer, 3 for example, return a string with a murmur: "1 sheep...2 sheep...3 sheep...".
Input will always be valid, i.e. no negative integers.
*
*/

func countSheep(num int) string {
	s := " sheep..."
	var s1 string
	for 1 <= num {
		s1 = strconv.Itoa(num) + s + s1
		num = num - 1
	}
	return s1
}
