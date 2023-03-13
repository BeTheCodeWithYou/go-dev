package main

import "fmt"

func main() {

	in := []int{7, 12, 13, 10, 3,7,9,24}
	in = filterOddNumber(in)
	fmt.Println(in)

}

func filterOddNumber(in []int) []int {
	odd := []int{}
	for i, v := range in {
		if v % 2 != 0 {
			odd = append(odd, in[i])
		}
	}
	return odd
}
