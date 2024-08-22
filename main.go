/*
Write a Go program that reads a string from the user and checks if it
 is a palindrome. A palindrome is a word, phrase, number, or other
 sequences of characters that read the same forward and backward
 (ignoring spaces, punctuation, and capitalization).*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Print("Enter a string: ")
	fmt.Scanln(&input)
	input = strings.ToLower(input)
	input = strings.ReplaceAll(input, " ", "")

	if isPlindrome(input) {
		fmt.Println(input, "is a palindrome")
	} else {
		fmt.Println(input, "is not a palindrome")
	}
}

func isPlindrome(s string) bool {
	for i := 0; i <= len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
