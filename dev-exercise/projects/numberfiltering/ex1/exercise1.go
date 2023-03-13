package main

import "fmt"

func main() {

	in := []int{1, 2, 3, 8, 9, 10, 17, 18, 33}
	in = filterEvenNum(in)
	fmt.Println(in)
}

func filterEvenNum(in []int) []int {
	evn := []int{}
	for i, v := range in {
		if (v % 2) == 0 {
			evn = append(evn, in[i])
		}
	}
	return evn
}
