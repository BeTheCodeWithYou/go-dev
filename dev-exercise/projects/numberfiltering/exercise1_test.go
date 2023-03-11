package main

import "testing"

func TestFilterEven(t *testing.T) {

	input := []int{1, 2}

	output := filterEven(input)

	for _, v := range output {
		if (v % 2) != 0 {
			t.Errorf("even number filtering failed for %v input ", input)
		}
	}

}
