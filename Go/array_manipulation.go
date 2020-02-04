package main

import "fmt"

// Complete the arrayManipulation function below.
func arrayManipulation(n int32, queries [][]int32) int64 {
	var i int32
	var a, b, k int64
	arr := make([]int64, n)

	var largest int64

	for _, element := range queries {
		a = int64(element[0] - 1)
		b = int64(element[1] - 1)
		k = int64(element[2])

		arr[a] += k
		if b+1 < int64(n) {
			arr[b+1] -= k
		}
	}

	largest = arr[0]
	for i = 1; i < n; i++ {
		arr[i] += arr[i-1]

		if largest < arr[i] {
			largest = arr[i]
		}
	}
	return int64(largest)
}

func main() {
	arr := [][]int32{
		{2, 6, 8},
		{3, 5, 7},
		{1, 8, 1},
		{5, 9, 15},
	}

	fmt.Println(arrayManipulation(10, arr))
}
