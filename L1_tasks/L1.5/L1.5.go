package main

import (
	"fmt"
	"time"
)

func main() {
	const N = 5
	dataChan := make(chan int)

	go func() {
		i := 0
		for {
			dataChan <- i
			i++
			time.Sleep(500 * time.Millisecond)
		}
	}()

	timeout := time.After(N * time.Second)

	fmt.Printf("Программа будет работать %d секунд...\n", N)

	for {
		select {
		case data := <-dataChan:
			fmt.Printf("Получено значение: %d\n", data)
		case <-timeout:
			fmt.Println("\nВремя вышло! Программа завершается.")
			return
		}
	}
}