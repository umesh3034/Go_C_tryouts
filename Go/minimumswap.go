package main

import "fmt"

// Complete the minimumSwaps function below.
func minimumSwaps(arr []int32) int32 {
	var swapcount, i, arrlen int32
	arrlen = int32(len(arr))
	for i = 0; i < arrlen; {
		fmt.Println("i :", i, arr[i])
		if arr[i] != i+1 {
			temp := arr[arr[i]-1]
			arr[arr[i]-1] = arr[i]
			arr[i] = temp
			swapcount++
		} else {
			i++
		}
	}
	return swapcount
}

func main() {
	fmt.Println(minimumSwaps([]int32{7, 1, 3, 2, 4, 5, 6}))
}
