package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("problems.csv")
	if err != nil {
		fmt.Printf("Error opening file %v", err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	s, err := r.ReadAll()
	if err != nil {
		fmt.Printf("error reading file data %v", err)
	}
	// fmt.Println(s)
	// q := s[1]
	// fmt.Println(q)
	// fmt.Println(q[0])
	// fmt.Print(q[1])
	var ua int
	var crt int
	var wng int
	for _, q := range s {
		fmt.Printf("What is %v ? ", q[0])
		fmt.Scanln(&ua)
		if err != nil {
			fmt.Printf("Invalid input %v", err)
		}
		a, err := strconv.Atoi(q[1])
		if err!=nil {
			fmt.Println("Error converting string to int ", err)
		}
		if ua == a {
			fmt.Println("Correct Answer!")
			crt++
		} else {
			fmt.Println("Wrong!")
			wng++
		}

	}
	fmt.Printf("Total Questions: %v \n", len(s))
	fmt.Printf("Correct: %v \n", crt)
	fmt.Printf("Wrong:  %v ", wng)
}
