package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func Intersection(a, b []int) []int {
	sort.Ints(a)
	sort.Ints(b)

	var result []int
	i, j := 0, 0

	for i < len(a) && j < len(b) {
		if a[i] == b[j] {
			result = append(result, a[i])
			i++
			j++
		} else if a[i] < b[j] {
			i++
		} else {
			j++
		}
	}
	return result
}

func main() {
	var a []int
	var b []int

	fmt.Println("Введите элементы первого среза A через пробел и нажмите Enter:")
	reader := bufio.NewReader(os.Stdin)
	inputA, _ := reader.ReadString('\n')
	numsA := strings.Split(strings.TrimSpace(inputA), " ")

	for _, s := range numsA {
		var num int
		_, err := fmt.Sscanf(s, "%d", &num)
		if err == nil {
			a = append(a, num)
		}
	}
	fmt.Println("Введите элементы второго среза B через пробел и нажмите Enter:")
	inputB, _ := reader.ReadString('\n')
	numsB := strings.Split(strings.TrimSpace(inputB), " ")

	for _, s := range numsB {
		var num int
		_, err := fmt.Sscanf(s, "%d", &num)
		if err == nil {
			b = append(b, num)
		}
	}

	result := Intersection(a, b)
	fmt.Printf("Пересечение: %v\n", result)
}
