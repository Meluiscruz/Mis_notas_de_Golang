package main

import (
	"fmt"
	"strings"
)

func isPalindrome(text string) bool {
	var textReverse string
	text = strings.ToLower(text)
	text = strings.ReplaceAll(" ", text, text)
	fmt.Println(text)

	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}

	if text == textReverse {
		return true
	} else {
		return false
	}

}

func main() {

	fmt.Println(isPalindrome("aNita la va la tina"))
}
