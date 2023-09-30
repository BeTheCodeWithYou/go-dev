package main

/**
Remove First and Last character

It's pretty straightforward. Your goal is to create a function that removes the first and last characters of a string.
You're given one parameter, the original string. You don't have to worry with strings with less than two characters.

**/

func removeChar(s string) string {
	return s[1 : len(s)-1] // this is possible because when strings are passed to a function, they are passed as slice.
}
