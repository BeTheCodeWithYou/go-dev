package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

	res, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("error while reading from http url", err)
		os.Exit(1)
	}

	// bs := make([]byte, 99999)
	// res.Body.Read(bs)
	// fmt.Println(string(bs))

	// OR as below

	io.Copy(os.Stdout, res.Body)
}
