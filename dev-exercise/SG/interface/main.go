package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

func main() {

	eg := englishBot{}
	//sb := spanishBot{}

	printGreeting(eg)
	//printGreeting(sb)
}

func (eb englishBot) getGreeting() string {
	return "hello"
}

func (sb spanishBot) getGreeting() string {
	return "Hola!"
}

/*func printGreeting(sb spanishBot) {
	fmt.Println(sb.getGreeting())
}*/

func printGreeting(eg englishBot){
	fmt.Println(eg.getGreeting())
}