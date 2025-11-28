package main

import (
	"fmt"
)

func ReverseString(input string) string {

	runes := []rune(input)

	i, j := 0, len(runes)-1

	for i < j {

		runes[i], runes[j] = runes[j], runes[i]

		i++
		j--
	}

	return string(runes)
}

func main() {
	str1 := "главрыба"
	fmt.Printf("%s |  %s\n", str1, ReverseString(str1))

	str2 := "Go рулит"
	fmt.Printf("%s | %s\n", str2, ReverseString(str2))
}
