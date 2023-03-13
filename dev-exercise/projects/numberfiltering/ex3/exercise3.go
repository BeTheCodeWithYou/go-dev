package main

import "fmt"

func main() {

	in := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 45}

	in = filterPrimeNumber(in)
	fmt.Println(in)

}

func filterPrimeNumber(in []int) []int {
	p := []int{}
	ts := make([]int, len(in))
	copy(ts, in)
	var ct int
	for _, v := range in {
		ct = 0
		if v > 1 {
			for _, tv := range ts {
				if tv>1 && v%tv == 0 {
					ct++
				}
			}
		}
		if ct == 1 {
			p = append(p, v)
		}
	}
	return p
}
