package main

import "fmt"

func mergeSort(a, b []int) []int {
	lenA, lenB := len(a), len(b)
	res := make([]int, 0, lenA+lenB)

	idxA, idxB := 0, 0
	for idxA+idxB < lenA+lenB {
		if idxB == lenB || (idxA < lenA && a[idxA] < b[idxB]) {
			res = append(res, a[idxA])
			idxA++
		} else {
			res = append(res, b[idxB])
			idxB++
		}

	}

	return res
}

func sort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	middle := len(arr) / 2
	left, right := arr[:middle], arr[middle:]

	left, right = sort(left), sort(right)

	return mergeSort(left, right)
}

func main() {
	fmt.Println(sort([]int{3, -1, 14, 2, 99, 17, 14, 105, 0, 1, 1, 45, 32, 88, 12, 85, 35, 66, 14}))
}
