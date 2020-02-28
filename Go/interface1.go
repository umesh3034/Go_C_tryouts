package main

import (
	"fmt"
)

//interface def
type VowelsFinder interface {
	FindVowels() []rune
}

type MyString string

//Mystring implements VowelsFinder
func (ms MyString) FindVowels() []rune {
	var vowels []rune
	for _, rune := range ms {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowels = append(vowels, rune)
		}
	}

	return vowels
}

func main() {
	name := MyString("sam anderson")
	var v VowelsFinder
	v = name
	fmt.Printf("vowels : %c", v.FindVowels())
}
