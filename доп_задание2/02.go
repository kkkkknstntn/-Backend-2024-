package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func countWords(filename string) map[string]int {
	wordCounts := make(map[string]int)
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return wordCounts
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		for _, word := range words {
			wordCounts[word]++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка при сканировании файла:", err)
	}

	return wordCounts
}

func main() {
	filename := "input.txt" // Имя файла, замените на имя вашего файла
	wordCounts := countWords(filename)

	// Сортируем слова по количеству вхождений в убывающем порядке
	sortedWords := make([]string, 0, len(wordCounts))
	for word := range wordCounts {
		sortedWords = append(sortedWords, word)
	}
	sort.Slice(sortedWords, func(i, j int) bool {
		return wordCounts[sortedWords[i]] > wordCounts[sortedWords[j]]
	})

	// Выводим список слов и их частоты
	for _, word := range sortedWords {
		fmt.Printf("%s: %d\n", word, wordCounts[word])
	}
}
