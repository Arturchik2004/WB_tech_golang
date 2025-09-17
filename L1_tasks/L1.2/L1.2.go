package main

import (
	"fmt"
	"sync"
)

func main() {
	data := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup

	for i, num := range data {
		wg.Add(1)
		go func(i int, num int) {
			defer wg.Done()
			n := num * num
			fmt.Printf("Квадрат числа %d равен %d\n", num, n)
			data[i] = n
		}(i, num)
	}

	wg.Wait()

	fmt.Printf("Новый, перезаписанный слайс: %v\n", data)
}
