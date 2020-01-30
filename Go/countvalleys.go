package main

import "fmt"

func countingValleys(n int32, s string) int32 {
	sealevel := true
	var countingValleys int32
	var up, down, i int32
	for i = 0; i < n; i++ {
		if sealevel == true {
			fmt.Printf("1111 %c\n", s[i])
			if s[i] == 'U' {
				up++
			} else {
				down++
			}
			sealevel = false
		} else if sealevel == false && up != 0 {
			fmt.Printf("2222 %c\n", s[i])
			if s[i] == 'U' {
				up++
			} else {
				up--
				if up == 0 {
					sealevel = true
				}
			}
		} else if sealevel == false && down != 0 {
			fmt.Printf("3333 %c\n", s[i])
			if s[i] == 'U' {
				down--
				if down == 0 {
					sealevel = true
					fmt.Println("increment count")
					countingValleys++
				}
			} else {
				down++
			}
		}
	}
	return countingValleys
}

func main() {
	fmt.Println(countingValleys(12, "DDUUDDUDUUUD"))
}
