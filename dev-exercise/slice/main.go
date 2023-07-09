package main

import "fmt"

func main() {

	s1 := make([]int, 2, 5) // lengh=2, capacity=5
	printSlice("s1", s1)

	s2 := s1

	// s2 slice will have exactly same elements, lenght=2 and same capacity=5
	//fmt.Printf("s1 slice: %v, s1 len: %v, s1 cap: %v \n", s1, len(s1), cap(s1))
	printSlice("s2", s2)

	// s2 will have new value at index 2, len=3, cap=5
	s2 = append(s2, 1)
	printSlice("s2", s2)

	// new value of s2 ( 1 which we just appended ) will NOT be visible to s1.
	// Though, both s1 and s2 oints to same underlaying array

	printSlice("s1", s1)

	s3 := s1[1:3]

	//now, s3 slice will be of len=2 ( 1st index included, and last index excluded),
	// cap=4 ( because index is starting from 1 and s1 original cap is 5, so 1,2,3,4 index makes the cap size 4 )
	printSlice("s3", s3)

}

func printSlice(sname string, s []int) {
	fmt.Printf("slice name: %v \n ele: %v, len: %v, cap: %v \n", sname, s, len(s), cap(s))
}
