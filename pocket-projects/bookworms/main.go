package main

import (
	"fmt"
	"os"
)

func main() {

	b, err := loadBookworms("testdata/bookworms.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to load bookworms: %s\n", err)
		os.Exit(1)
	}
	fmt.Println(b)
}
