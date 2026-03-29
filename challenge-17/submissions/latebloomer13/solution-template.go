package main

import (
	"fmt"

)

func main() {
	// Get input from the user
	var input string
	fmt.Print("Enter a string to check if it's a palindrome: ")
	fmt.Scanln(&input)

	// Call the IsPalindrome function and print the result
	result := IsPalindrome(input)
	if result {
		fmt.Println("The string is a palindrome.")
	} else {
		fmt.Println("The string is not a palindrome.")
	}
}

// IsPalindrome checks if a string is a palindrome.
// A palindrome reads the same backward as forward, ignoring case, spaces, and punctuation.
func IsPalindrome(s string) bool {
	
	clean_s := CleanString(s)
	
	if len(clean_s) == 0 {
	    return true
	}
	
	runes := []rune(clean_s)
	fmt.Println(runes)
	for i := 0; i <= len(runes) / 2; i++ {
	    if runes[i] != runes[len(runes)-i-1] {
	        return false
	    }
	}
	
	return true
}

func CleanString(s string) string {
    clean_s := ""
	
	for _, char := range s {
	   	if char >= 'A' && char <= 'Z' {
	        clean_s = clean_s + string(char + ('a' - 'A'))
	    } else {
	        switch char {
	            case ' ', ',', '!', '?', '.', '\'', ':', '@', '#', '$', '%', '^', '&', '*', '(', ')', '"':
	        default: clean_s += string(char)
	   }
	}
	    
	}
	
	return clean_s
}
