package main

import (
	"fmt"
)

func maxSum(numbers []int) int {
	// Находим два наибольших числа
	max1, max2 := numbers[0], numbers[0]
	for _, num := range numbers {
		if num > max1 {
			max2 = max1
			max1 = num
		} else if num > max2 {
			max2 = num
		}
	}

	// Возвращаем сумму двух наибольших чисел
	return max1 + max2
}

func main() {
	numbers := []int{2, 7, 4, 1, 8}
	fmt.Println(maxSum(numbers))
}
