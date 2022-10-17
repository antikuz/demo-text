package main

import (
	"fmt"
	"strings"
)


func countChars(text string) {
	alphabetENG := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	alphabetRUS := "АБВГДЕЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯабвгдежзийклмнопрстуфхцчшщъыьэюя"
	result := ""
	
	for _, char := range alphabetENG {
		count := strings.Count(text, string(char))
		if count > 0 {
			result += fmt.Sprintf("%c%d", char, count)
		}
	}

	for _, char := range alphabetRUS {
		count := strings.Count(text, string(char))
		if count > 0 {
			result += fmt.Sprintf("%c%d", char, count)
		}
	}

	fmt.Printf("Answer: %s", result)
}


func main() {
	var text string
    
	fmt.Print("Enter chars sequence: ")
    fmt.Scanln(&text)
	countChars(text)
}
