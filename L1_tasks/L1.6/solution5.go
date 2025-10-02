package main

import (
	"fmt"
	"os"
	"time"
)

func worker() {
	defer fmt.Println("Этот defer НЕ будет выполнен.")
	time.Sleep(2 * time.Second)
	fmt.Println("Горутина завершила работу (не успеет).")
}

func main() {
	go worker()
	fmt.Println("Программа завершится через 1 секунду.")
	time.Sleep(1 * time.Second)
	os.Exit(1)
}
