package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
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
	var crt, wrng int
	var ua string
	for _, q := range s {
		fmt.Printf("What is %v ? ", q[0])
		fmt.Scanln(&ua)
		if err != nil {
			fmt.Printf("Invalid input %v", err)
		}
		if ua == "" || ua != q[1]{
			wrng++
		} else {
			crt++
		}

	}
	fmt.Printf("Total Questions: %v \n", len(s))
	fmt.Printf("Correct: %v \n", crt)
	fmt.Printf("Wrong:  %v ", wrng)
}
