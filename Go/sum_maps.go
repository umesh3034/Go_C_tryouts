package main

import "fmt"

func twoSum(nums []int, sum int) [][]int {
	var pairs [][]int = make([][]int, 0)
	var hashmap map[int]int = make(map[int]int, 0)
	for index, value := range nums {
		diff := sum - value
		fmt.Println("11 -", diff, sum, "-", value, ":", index)
		if hashentry, ok := hashmap[diff]; ok {
			fmt.Println("22 -", hashentry, ok)
			if hashentry != index {
				pairs = append(pairs, []int{diff, value})
				fmt.Println("33 -", pairs)
			}
		} else {
			fmt.Println("44 -", value, index)
			hashmap[value] = index
		}
	}
	fmt.Println("hashmap -", hashmap)
	return pairs
}

func main() {
	fmt.Println(twoSum([]int{1, 4, 2, 3, 6, 5, 0, 3}, 6))
}
