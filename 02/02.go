package main

import (
	"fmt"
	"sort"
)

func sortAbyB(A []int, B []int) []int {
	indexMap := make(map[int]int)
	// сохдаем map с ключами в виде элеменетов массива B
	for _, b := range B {
		indexMap[b] = 0
	}
	var result1 []int
	var result2 []int
	// Считаем количество вхождений каждого элемента массива B в массив A
	// Параллельно добавляем все остальные элементы в массив result2
	for _, a := range A {
		if _, ok := indexMap[a]; ok {
			indexMap[a] += 1
		} else {
			result2 = append(result2, a)
		}
	}
	//Добавляем в result1 элементы в порядке массива B
	for _, a := range B {
		for i := 0; i < indexMap[a]; i++ {
			result1 = append(result1, a)
		}
	}
	//Сливаем result1 и result2
	sort.Slice(result2, func(i, j int) bool {
		return result2[j] < result2[i]
	})
	result1 = append(result1, result2...)

	return result1
}

func main() {
	A := []int{2, 4, 1, 3, 2, 4, 6, 7, 9, 2, 19}
	B := []int{2, 1, 4, 3, 6, 9}
	fmt.Println(sortAbyB(A, B))
}
