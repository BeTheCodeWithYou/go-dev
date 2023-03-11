package main

import "fmt"

/*
Given a list of integers, write a program to return only the even numbers from this list.

Sample Input: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10
Sample Output: 2, 4, 6, 8, 10

*/
func main() {

	in := []int{1, 2, 3, 8, 9, 10, 17, 18, 33}
	in = filterEven(in)
	fmt.Println(in)
}

func filterEven(in []int) []int {

	for i, v := range in {
		if (v % 2) != 0 {
			copy(in[i:], in[i+1:]) // shift in[i+1:] left one index
			in[len(in)-1] = 0      // erase last element
			in = in[:len(in)-1]    // turncate slice
		}
	}
	return in
}
