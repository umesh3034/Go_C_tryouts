package main

import "fmt"

func swap(a []int32, i int, j int) {
	temp := a[i]
	a[i] = a[j]
	a[j] = temp
}

// Complete the maximumToys function below.
func maximumToys(prices []int32, k int32) int32 {
	n := len(prices)
	var sum int32
	for i := 0; i < n; i++ {

		for j := n - 1; j > i; j-- {
			// Swap adjacent elements if they are in decreasing order
			if prices[j] < prices[j-1] {
				swap(prices, j, j-1)
			}
		}
		fmt.Println(prices)
		if sum+prices[i] > k {
			return int32(i)
		} else {
			sum += prices[i]
		}
	}
	return 0
}

func main() {

	arr := []int32{33324560, 77661073, 31948330, 21522343, 97176507, 5724692, 24699815, 12079402,
		6479353, 28430129, 42427721, 57127004, 26256001, 29446837, 65107604, 9809008, 65846182,
		8470661, 13597655, 360}
	fmt.Println(arr)
	fmt.Println(maximumToys(arr, 10000))
}
