package main

import "fmt"

func main() {

	v := 1
	incr(&v)
	fmt.Print(incr(&v))

}

func incr(p *int) int {
	*p++
	return *p
}
