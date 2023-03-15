package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	in, opt := userInputProcessing()
	switch opt {
	case 1:
		filterEvenNumber(in)
	case 2:
		filterOddNumber(in)
	case 3:
		filterPrimeNumber(in)
	}
}

func userInputProcessing() ([]int, int) {

	fmt.Println("Please choose an option for number filtering program.")
	fmt.Println("1. Even numbers")
	fmt.Println("2. Odd Numbers")
	fmt.Println("3. Prime Numbers")
	var opt int
	fmt.Scanln(&opt)
	fmt.Println("Provide list of input numbers with comma seperated values")
	scanner := bufio.NewScanner(os.Stdin)
	var str string
	if scanner.Scan() {
		str = scanner.Text()
	}

	input := strings.Split(str, ",")
	var in []int
	for _, v := range input {
		val, _ := strconv.Atoi(v)
		in = append(in, val)
	}
	return in, opt
}

func filterEvenNumber(in []int) {
	evn := []int{}
	for i, v := range in {
		if (v % 2) == 0 {
			evn = append(evn, in[i])
		}
	}
	fmt.Println("Even number ->", evn)
}

func filterOddNumber(in []int) {
	odd := []int{}
	for i, v := range in {
		if v%2 != 0 {
			odd = append(odd, in[i])
		}
	}
	fmt.Println("Odd number ->", odd)
}

func filterPrimeNumber(in []int) {
	p := []int{}
	ts := make([]int, len(in))
	copy(ts, in)
	var ct int
	for _, v := range in {
		ct = 0
		if v > 1 {
			for _, tv := range ts {
				if tv > 1 && v%tv == 0 {
					ct++
				}
			}
		}
		if ct == 1 {
			p = append(p, v)
		}
	}
	fmt.Println("Prime number ->", p)
}