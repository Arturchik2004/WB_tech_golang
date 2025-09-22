package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(id int, wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
	fmt.Printf("Воркер %d: запущен\n", id)
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Воркер %d: получил сигнал завершения. Выхожу.\n", id)
			return
		default:
			fmt.Printf("Воркер %d: работаю...\n", id)
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	fmt.Println("Программа запущена. Нажмите Ctrl+C для завершения.")

	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	numWorkers := 5
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, &wg, ctx)
	}

	<-sigChan
	fmt.Println("\nПолучен сигнал прерывания. Инициирую завершение...")

	cancel()
	wg.Wait()

	fmt.Println("Все воркеры завершили работу. Программа выходит.")
}
