package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {

	//var f = flag.String("filename", "problems.csv", "File containing questions list")
	var t = flag.Int("timer", 30, "question timer")
	flag.Parse()
	fmt.Println("timer flag->", *t)
	t1 := time.NewTimer(30 * time.Second)
	<- t1.C
	fmt.Println("timer ended")


}
