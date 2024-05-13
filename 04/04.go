package main

import (
	"fmt"
)

func binary_search(array []int, to_search int) int {
	low := 0
	high := len(array) - 1
	for low <= high {
		mid := (low + high) / 2 // Середина массива
		if array[mid] == to_search {
			return mid // Элемент найден, возвращаем индекс
		}
		if array[mid] > to_search { //Элемент меньше середины
			high = mid - 1
		} else { //Элемент меньше середины
			low = mid + 1
		}
	}
	return -1 // Элемент не найден
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 5

	index := binary_search(arr, target)
	if index != -1 {
		fmt.Printf("Элемент %d найден на позиции %d\n", target, index)
	} else {
		fmt.Printf("Элемент %d не найден\n", target)
	}
}
