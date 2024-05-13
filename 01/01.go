package main

import (
	"fmt"
)

// решил реализовать сортировку с помощью quicksort'а
func quickSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	left, right := 0, len(arr)-1

	pivot := arr[right]

	for i := range arr {
		if arr[i] < pivot {
			arr[left], arr[i] = arr[i], arr[left]
			left++
		}
	}

	arr[left], arr[right] = arr[right], arr[left]

	quickSort(arr[:left])
	quickSort(arr[left+1:])
}

func main() {
	numbers := []int{10, 7, 8, 9, 1, 5}
	fmt.Println("Неотсортированный:", numbers)

	quickSort(numbers)
	fmt.Println("Отсортированный:", numbers)
}
