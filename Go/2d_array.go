package main

import "fmt"

// Complete the hourglassSum function below.
func hourglassSum(arr [][]int32) int32 {
	var maxhourglass int32 = -2147483648
	var temp int32 = 0
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			temp = arr[i][j] + arr[i][j+1] + arr[i][j+2] + arr[i+1][j+1] + arr[i+2][j] + arr[i+2][j+1] + arr[i+2][j+2]
			if temp > maxhourglass {
				maxhourglass = temp
			}
		}
	}
	return maxhourglass
}

func main() {
	array := [][]int32{
		{-1, -1, 0, -9, -2, -2},
		{-2, -1, -6, -8, -2, -5},
		{-1, -1, -1, -2, -3, -4},
		{-1, -9, -2, -4, -4, -5},
		{-7, -3, -3, -2, -9, -9},
		{-1, -3, -1, -2, -4, -5},
	}

	fmt.Println(hourglassSum(array))
}
