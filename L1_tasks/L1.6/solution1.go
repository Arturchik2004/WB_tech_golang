package main

import (
	"fmt"
	"sync"
	"time"
)

func cat(cats <-chan string, wg *sync.WaitGroup) {

	defer wg.Done()
	for meow := range cats {
		fmt.Println(meow)
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println("Проснулся злой сосед, прогнал всех котов с крыши и закрыл канал!")
}

func main() {
	cats := make(chan string, 5)
	var wg sync.WaitGroup

	wg.Add(1)
	go cat(cats, &wg)

	for i := 1; i <= 5; i++ {
		cats <- fmt.Sprintf("Кот %d говорит: Мяу!", i)
	}
	close(cats)
	wg.Wait()

	fmt.Println("Все коты ушли, на крыше тихо.")
}
