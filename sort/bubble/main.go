package main

import "fmt"

func sort(arr []int) []int {
	n := len(arr)
	if n <= 1 {
		return arr
	}

	swapped := true
	for swapped {
		swapped = false
		for i := 1; i < n; i++ {
			if arr[i-1] > arr[i] {
				swap(arr, i-1, i)
				swapped = true
			}
		}
	}

	return arr
}

func swap(arr []int, i, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}

func main() {
	fmt.Println(sort([]int{3, -1, 14, 2, 99, 17, 14, 105, 0, 1, 1, 45, 32, 88, 12, 85, 35, 66, 14}))
}
