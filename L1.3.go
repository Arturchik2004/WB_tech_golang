package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Воркер %d задачу обработал %d\n", id, job)
		time.Sleep(time.Second)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Использование: go run L1.3.go N \n, где N - количество воркеров > 0\n")
		return
	}

	numWorkers, err := strconv.Atoi(os.Args[1])
	if err != nil || numWorkers <= 0 {
		fmt.Println("Ошибка преобразования числа:", err)
		return
	}

	var wg sync.WaitGroup
	jobs := make(chan int)

	fmt.Printf("Запускаем %d воркеров\n", numWorkers)

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	// Отслеживание сигнала завершения (Ctrl+C)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		jobID := 1
		for {
			select {
			case <-sigChan:
				fmt.Println("\nПолучен сигнал завершения. Закрываем канал...")
				close(jobs)
				return
			default:
				jobs <- jobID
				jobID++
				time.Sleep(time.Second)
			}
		}
	}()

	wg.Wait()

	fmt.Println("Все воркеры завершили работу")
}
