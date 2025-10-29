package main

import "fmt"

func unic(sequence []string) []string {
	set := make(map[string]struct{})

	for _, item := range sequence {
		set[item] = struct{}{}
	}

	uniq := make([]string, 0, len(set))
	for key := range set {
		uniq = append(uniq, key)
	}

	return uniq
}

func main() {

	sequence := []string{"cat", "cat", "dog", "cat", "tree"}
	uniq := unic(sequence)
	fmt.Println("Исходная последовательность:", sequence)
	fmt.Println("Полученное множество (в виде среза):", uniq)

}
