package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {

	var fn = flag.String("filename", "problems.csv", "file name containing questions")
	flag.Parse()
	fmt.Println(*fn)
	f, err := os.Open(*fn)
	if err != nil {
		fmt.Printf("Error opening file %v", err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	s, err := r.ReadAll()
	if err != nil {
		fmt.Printf("error reading file data %v", err)
	}
	var ua, crt, wrng int
	for _, q := range s {
		fmt.Printf("What is %v ? ", q[0])
		fmt.Scanln(&ua)
		if err != nil {
			fmt.Printf("Invalid input %v", err)
		}
		a, err := strconv.Atoi(q[1])
		if err != nil {
			fmt.Println("Error converting string to int ", err)
		}
		if ua == a {
			crt++
		} else {
			wrng++
		}

	}
	fmt.Printf("Total Questions: %v \n", len(s))
	fmt.Printf("Correct: %v \n", crt)
	fmt.Printf("Wrong:  %v ", wrng)
}
