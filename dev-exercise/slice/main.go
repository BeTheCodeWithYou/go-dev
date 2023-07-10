package main

import "fmt"

func main() {

	s1 := make([]int, 2, 5) // lengh=2, capacity=5
	// slice 1 initialized with zeroed value as slice being integer array.
	printSlice("s1", s1)

	// new element will get append next to the default value
	s1 = append(s1, 7)

	printSlice("s1", s1)
	
	// s2 slice will have exactly same elements, lenght=3 and same capacity=5 as of s1
	s2 := s1
	printSlice("s2", s2)

	s3 := s1[2:4]

	//now, s3 slice will be of len=2 ( 1st index included, and last index excluded),
	// cap=3 ( because index is starting from 2 and s1 original cap is 5, so 2,3,4 index makes the cap size 3 )
	printSlice("s3", s3)

	s3 = append(s3,8)
	s3 = append(s3, 9)

	printSlice("s3", s3)

	s2 = append(s2, 1)

	printSlice("s2", s2)
	printSlice("s1", s1)

}

func printSlice(sname string, s []int) {
	fmt.Printf("slice name: %v \n ele: %v, len: %v, cap: %v \n", sname, s, len(s), cap(s))
}
