<h1 align = "center"><img width="683" height="415" alt="image" src="https://github.com/user-attachments/assets/657fa81f-53ea-435e-aa65-4ea30e686606" />
</h1>

## Описание программы 

### Импорты
*fmt* - для вывода и *time* для задания времени
 ```go
 package main

import (
	"fmt"
	"time"
)
 ```
### Основная функция *main*
В основной функции задаем параметр N (время, поистечении которого программма завершится) `const N = 5` и создаем канал для передачи числовых значений `dataChan := make(chan int)`.

Затем запускаем горутину, которая работает в бесконечном цикле. 
```go
go func() {
		i := 0
		for {
			dataChan <- i
			i++
			time.Sleep(500 * time.Millisecond)
		}
	}()
```
Главный цикл с select читает данные из канала пока не выйдет время `case <-timeout`
```go
	for {
		select {
		case data := <-dataChan:
			fmt.Printf("Получено значение: %d\n", data)
		case <-timeout:
			fmt.Println("\nВремя вышло! Программа завершается.")
			return
		}
	}
```
## Запуск программы
```bash
go run L1.5.go
```
<img width="354" height="179" alt="image" src="https://github.com/user-attachments/assets/71449d1b-bbb8-4bf5-83a7-7f8bc78ff54d" />


## Весь код
```go
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
```



