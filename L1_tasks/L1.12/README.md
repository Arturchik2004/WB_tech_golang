# L1.12 - Собственное множество строк
<img width="828" height="511" alt="image" src="https://github.com/user-attachments/assets/37fbbc03-29f7-4bbc-96fe-97a7dc5f0043" />


## Описание

В этом задании мы реализуем собственное множество строк на языке Go. Мы будем использовать мапу (map) для хранения уникальных строк из исходной последовательности. Затем мы извлечем ключи из мапы и представим их в виде среза (slice).

## Алгоритм

1. Создать пустую мапу для хранения уникальных строк.
```go
    set := make(map[string]struct{})
```
2. Пройти по входному срезу и для каждого элемента добавить ключ в мапу: `set[item] = struct{}{}`.
```go
	for _, item := range sequence {
		set[item] = struct{}{}
	}
``` 
3. Создать результирующий срез `uniq` с capacity равной `len(set)` для экономии аллокаций.
```go
	uniq := make([]string, 0, len(set))
```
4. Пройти по ключам мапы и добавить их в срез `uniq` (порядок не гарантируется).
```go
	for key := range set {
		uniq = append(uniq, key)
	}
```
5. Вернуть `uniq`.
go```
	return uniq
```

## Как запустить
```bash
go run L1.12.go
```

## Вывод
```bash
Исходная последовательность: [cat cat dog cat tree]
Полученное множество (в виде среза): [cat dog tree]
```

## Также
В этом задании можно добавить ввод строк через пробел с консоли. Для этого в func main() вместо заданного среза `sequence := []string{"cat", "cat", "dog", "cat", "tree"}` нужно использовать следующий код:
```go
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("Введите строки через пробел:")
    input, _ := reader.ReadString('\n')
    input = strings.TrimSpace(input)
    sequence := strings.Split(input, " ")
```
И подключить библиотеки `bufio` и `strings`:
```go
import (
    "bufio"
    "fmt"
    "os"
    "strings"
)
```
