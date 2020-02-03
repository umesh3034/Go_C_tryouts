package main

import "fmt"

func swap(arr []int32, i int32, j int32) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}

func partition(arr []int32, low int32, high int32) int32 {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			swap(arr, i, j)
		}
	}
	swap(arr, i+1, high)
	return i + 1
}

func quicksort(arr []int32, low int32, high int32) {
	if low < high {
		pi := partition(arr, low, high)
		quicksort(arr, low, pi-1)
		quicksort(arr, pi+1, high)
	}
}

// Complete the minimumSwaps function below.
func sort(arr []int32) {
	arrlen := len(arr)
	quicksort(arr, 0, int32(arrlen-1))
}

func main() {
	arr := []int32{1, 6, 7, 3, 5, 4, 10, 8, 9, 2}
	fmt.Println("before : ", arr)
	sort(arr)
	fmt.Println("after : ", arr)
}
