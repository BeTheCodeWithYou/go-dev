package main

import "fmt"

type bot interface{
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {

	eg := englishBot{}
	sb := spanishBot{}

	printGreeting(eg)
	printGreeting(sb)
	//printGreeting(sb)
}

func (eb englishBot) getGreeting() string {
	return "hello"
}

func (sb spanishBot) getGreeting() string {
	return "Hola!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// func printGreeting(eg englishBot){
// 	fmt.Println(eg.getGreeting())
// }