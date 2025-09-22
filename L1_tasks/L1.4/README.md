# <p align="center"><img width="638" height="377" alt="image" src="https://github.com/user-attachments/assets/14e1803b-d8ba-42cc-8e72-6fe360099af3" /></p>

Для корректного завершения работы воркеров при получении сигнала SIGINT (Ctrl+C), я использую context.
Этот способ предпочтителен, так как:

 - Context.Context — это стандартный и общепринятый способ для управления отменой операций, установкой дедлайнов и передачей значений в рамках одного запроса или задачи в Golang.
 - Context позволяет отменить операции по нескольким причинам, а не только по сигналу от ОС. Например, можно легко добавить отмену по тайм-ауту (context.WithTimeout) или отмену родительской операции.
 - Использование context.Context явно делает код более читаемым и предсказуемым.
 ## Пример программы
 ```go
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
```
### Воркер
Основная горутина
Это функция, которая выполняет какую-то задачу в отдельной горутине. Select позволяет воркеру одновременно и работать, и слушать сигнал о завершении.

```go
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
```
### *main*
В *main* мы подготавливаем три ключевых компонента:
1. *context.WithCancel:* Создаёт контекст и функцию cancel. Вызов этой функции отправит сигнал завершения всем горутинам, которые слушают этот контекст.

2. *sync.WaitGroup:* Счётчик, который позволяет главной горутине дождаться выполнения всех остальных.

3. *signal.Notify:* Направляет сигналы операционной системы (как SIGINT от Ctrl+C) в наш канал sigChan.
```go
func main() {
	fmt.Println("Программа запущена. Нажмите Ctrl+C для завершения.")

	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

```
### Запуск горутин
Запускаем горутины в цикле
```go
numWorkers := 5
for i := 1; i <= numWorkers; i++ {
	wg.Add(1)
	go worker(i, &wg, ctx)
}
```
 ### Финальная часть программы. 
 Главная горутина блокируется и ждёт сигнала. Как только он получен, запускается процесс завершения: вызывается cancel() и программа ждёт, пока все воркеры не закончат свою работу.
 ```go
<-sigChan
fmt.Println("\nПолучен сигнал прерывания. Инициирую завершение...")

cancel()

wg.Wait()

fmt.Println("Все воркеры завершили работу. Программа выходит.")
}
 ```
 Весь код
 ```go
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
 ```
