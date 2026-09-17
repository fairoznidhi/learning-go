package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func ReadWords(filename string) ([]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return strings.Fields(string(data)), nil
}
func cleanWord(word string) string {
	word = strings.ToLower(word)
	word = strings.TrimFunc(word, unicode.IsPunct)
	return word
}

func CleanWords(words []string) []string {
	cleaned := make([]string, 0, len(words))
	for _, word := range words {
		cleaned = append(cleaned, cleanWord(word))
	}
	return cleaned
}

func UniqueWords(words []string) []string {
	uniqueMap := make(map[string]struct{})
	for _, word := range words {
		uniqueMap[word] = struct{}{}
	}
	uniqueWords := make([]string, 0, len(uniqueMap))
	for word := range uniqueMap {
		uniqueWords = append(uniqueWords, word)
	}
	return uniqueWords
}

func main() {
	words, err := ReadWords("sample.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	words = CleanWords(words)
	words = UniqueWords(words)
	fmt.Println(words)
}
