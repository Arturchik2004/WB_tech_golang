package main

import (
	"fmt"
)

func main() {
	input := "snow dog sun"
	fmt.Printf("Входная строка: %s\n", input)

	result := reverseWords(input)
	fmt.Printf("Выход: %s\n", result)
}

func reverseWords(s string) string {
	runes := []rune(s)
	reverse(runes, 0, len(runes)-1)

	start := 0
	for i := 0; i <= len(runes); i++ {
		if i == len(runes) || runes[i] == ' ' {
			reverse(runes, start, i-1)
			start = i + 1
		}
	}

	return string(runes)
}

func reverse(r []rune, start, end int) {
	for start < end {
		r[start], r[end] = r[end], r[start]
		start++
		end--
	}
}
