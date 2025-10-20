# L1.9 Конвейер чисел на Go
<img width="805" height="495" alt="image" src="https://github.com/user-attachments/assets/998e9397-4016-410f-b1ba-aec415b12710" />


Этот код демонстрирует создание простого конвейера с использованием горутин и каналов в языке Go.

## Описание

Программа выполняет следующие действия:
1.  Читает последовательность чисел, введенных пользователем в одну строку через пробел.

```go
    fmt.Println("Введите числа через пробел и нажмите Enter:")
	var m []int

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	nums := strings.Split(strings.TrimSpace(input), " ")

	for _, s := range nums {
		if num, err := strconv.Atoi(s); err == nil {
			m = append(m, num)
		}
	}

```


2.  Создает конвейер из двух этапов (двух горутин):
   
    - Горутина читает числа из исходного массива и отправляет их в первый канал.
      ```go 
    	func res(elem <-chan int, result chan<- int) {
	   	 	for val := range elem {
		    	result <- val * 2
	    	}
	    	close(result)
    	}

    	func main(){
        	...
        	for _, val := range m {
		    	elem <- val
	    	}
        	close(elem)
    	} 
		```
    - Вторая горутина читает числа из первого канала, умножает каждое на 2 и выводит результат в консоль.
	
		```go
    
    	...
    	func res2(results <-chan int, wg *sync.WaitGroup) {
	    	defer wg.Done()
	    	for val := range results {
		    	fmt.Println(val * val)
	    	}
    	}
    	...
		```
    

4.  Программа корректно ожидает завершения всех операций перед выходом, используя `sync.WaitGroup`.

    ```go
    var wg sync.WaitGroup
    wg.Add(1)
    go res2(results, &wg)
    wg.Wait()
    ```


## Как запустить

1.  Убедитесь, что у вас установлен Go.
2.  Перейдите в директорию с файлом:
    ```bash
    cd L1_tasks/L1.9
    ```
3.  Запустите программу:
    ```bash
    go run L1.9.go
    ```
4.  Введите числа через пробел и нажмите Enter.

## Пример работы

**Ввод:**
```1 2 3 4 5```

**Вывод:**
```
1
4
9
16
25 
```
