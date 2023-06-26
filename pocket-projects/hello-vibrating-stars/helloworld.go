package main

import (
	"flag"
	"fmt"
)

func main() {
	var lang string
	flag.StringVar(&lang, "lang", "en", "the required language e.g en, fr...")
	flag.Parse()

	greeting := greet(language(lang))
	fmt.Println(greeting)
}

type language string

var phraseBook = map[language]string{
	"en": "hello world",
	"fr": "Bonjour le monde",  // French
	"he": " שלום עולם ",       // Hebrew
	"vi": "Xin chào Thế Giới", // Vietnamese
	"hi": "हेलो विश्व",        // hindi
}

func greet(l language) string {

	language, ok := phraseBook[l]
	if !ok {
		return fmt.Sprintf("unsupported language: %q ", l)
	}
	return language

}
