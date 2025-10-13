package main

import "fmt"

func setBit(num int64, i uint, value int) int64 {
	if value == 1 {

		return num | (1 << i)
	}

	return num &^ (1 << i)
}

func main() {
	var number int64 = 5
	var position uint = 1
	var newValue = 0

	fmt.Printf("Исходное число: %d (%04b)\n", number, number)

	result := setBit(number, position, newValue)

	fmt.Printf("Устанавливаем бит на позиции %d в значение %d\n", position, newValue)
	fmt.Printf("Результат: %d (%04b)\n", result, result)
	fmt.Println("---")

	number = 9
	position = 3
	newValue = 1

	fmt.Printf("Исходное число: %d (%04b)\n", number, number)

	result = setBit(number, position, newValue)

	fmt.Printf("Устанавливаем бит на позиции %d в значение %d\n", position, newValue)
	fmt.Printf("Результат: %d (%04b)\n", result, result)
}
