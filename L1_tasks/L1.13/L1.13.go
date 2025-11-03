package main

import (
	"fmt"
)

func XOR(a int, b int) (int, int) {
	a ^= b
	b ^= a
	a ^= b
	return a, b
}

func SummSum(a int, b int) (int, int) {
	a = a + b
	b = a - b
	a = a - b
	return a, b
}

func main() {
	var a, b int
	var ch int
	fmt.Print("Выберите метод обмена значениями переменных:\n1 - с помощью XOR\n2 - с помощью сложения и вычитания\n")
	fmt.Scan(&ch)

	switch ch {
	case 1:
		fmt.Println("Введите два целых числа через пробел:")
		fmt.Scan(&a, &b)
		a, b = XOR(a, b)
		fmt.Println(a, b)
	case 2:
		fmt.Println("Введите два целых числа через пробел:")
		fmt.Scan(&a, &b)
		a, b = SummSum(a, b)
		fmt.Println(a, b)
	default:
		fmt.Println("Некорректный ввод")
	}
}
