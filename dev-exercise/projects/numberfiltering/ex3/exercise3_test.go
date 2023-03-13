package main

import (
	"fmt"
	"os"
	"testing"
)

func TestFilterPrimeNum(t *testing.T) {

	in := []int{1, 2, 3, 45, 67, 33, -2}
	ou := filterPrimeNumber(in)
	fmt.Println("Prime numbers are -", ou)

	for _, v := range ou {
		if v <= 1 {
			t.Errorf("%v is not a prime number", v)
			os.Exit(1)
		}
		if v != 2 && v%2 == 0 {
			t.Errorf("%v is not a prime number", v)
		}
	}
}
