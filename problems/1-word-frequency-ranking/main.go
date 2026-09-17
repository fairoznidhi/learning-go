// Write a Go program that takes a block of text and produces a ranked list of the most frequently occurring words.
package main

import (
	"fmt"
)

type WordCount struct {
	word  string
	count int
}

var punctuation = []byte{'.', ',', '!', '?', '"', '\'', ';', ':'}
var stopWords = []string{"the", "a", "an", "is", "it", "and", "of", "to"}

func toLower(text string) string {
	bs := []byte(text)
	for i, c := range bs {
		// fmt.Println(i, " ", string(c))
		if c >= 'A' && c <= 'Z' {
			bs[i] = c + 32
		}
	}
	return string(bs)
}

func isPunctuation(character byte) bool {
	for _, p := range punctuation {
		if character == p {
			return true
		}
	}
	return false
}

func extractWords(text string) []string {
	var words []string
	currentWord := ""
	for _, c := range text {
		if c == ' ' || c == '\n' || c == '\t' {
			if currentWord != "" {
				words = append(words, currentWord)
				currentWord = ""
			}
		} else {
			currentWord += string(c)
		}
	}
	if currentWord != "" {
		words = append(words, currentWord)
	}
	return words
}

func NormalizeWords(text string) []string {
	lower := toLower(text)
	bs := []byte(lower)

	removedPunc := make([]byte, 0, len(bs))
	for _, c := range bs {
		if !isPunctuation(c) {
			removedPunc = append(removedPunc, c)
		}
	}
	words := extractWords(string(removedPunc))
	return words
}

func WordFrequency(text string, topN int) []WordCount {
	normalized := NormalizeWords(text)
	wordMap := make(map[string]int)
	for _, word := range normalized {
		wordMap[word]++
	}

	counts := make([]WordCount, 0, len(wordMap))
	for word, count := range wordMap {
		counts = append(counts, WordCount{word: word, count: count})
	}
	return []WordCount{}
}

func main() {
	text := "The quick brown fox jumps over the lazy dog. The dog barks, and the fox runs."
	// result := WordFrequency(text, 5)
	result := NormalizeWords(text)
	fmt.Println(result)
}
