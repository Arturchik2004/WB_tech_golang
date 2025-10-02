package main

import (
	"fmt"
	"time"
)

func student(done <-chan bool) {
	for {
		select {
		case <-done:
			fmt.Println("Получен сигнал, студент уснул...")
			return
		default:
			fmt.Println("Студент пишет лекцию...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	done := make(chan bool)

	go student(done)

	time.Sleep(2 * time.Second)
	done <- true

	time.Sleep(500 * time.Millisecond)
}
