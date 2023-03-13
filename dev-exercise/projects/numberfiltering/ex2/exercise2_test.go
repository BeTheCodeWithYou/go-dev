package main

import "testing"

func TestFilterOddNum(t *testing.T) {

	i := []int{7, 8, 12}
	o := filterOddNumber(i)

	for _, v := range o {
		if v%2 == 0 {
			t.Errorf("odd number filtering failed for v %d", v)
		}
	}

}
