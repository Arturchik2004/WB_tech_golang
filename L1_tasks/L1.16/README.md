<img width="856" height="508" alt="image" src="https://github.com/user-attachments/assets/7cdac070-a003-4e37-b3e4-fa68e9e05411" />



# Быстрая сортировка (quicksort)


## Функция quicksort
quickSort реализует алгоритм быстрой сортировки.


Базовый случай рекурсии: если в массиве 0 или 1 элемент, он уже отсортирован.
```go
    if len(arr) < 2 {
        return
    }
```

Инициализируем указатели на начало и конец массива.
```go
    left, right := 0, len(arr)-1
```

Выбираем опорный элемент (pivot). 
В данном случае, это случайный элемент. Это помогает избежать худшего случая O(n^2) на уже отсортированных или обратно отсортированных данных.
```go
    rand.Seed(time.Now().UnixNano())
    pivotIndex := rand.Intn(len(arr))
    pivot := arr[pivotIndex]
```

Меняем опорный элемент с последним, чтобы он не мешал разделению.
```go
    arr[pivotIndex], arr[right] = arr[right], arr[pivotIndex]
```

Перемещаем все элементы, которые меньше опорного, в левую часть.
```go
    for i := range arr {
        if arr[i] < pivot {
            arr[left], arr[i] = arr[i], arr[left]
            left++
        }
    }
```

Возвращаем опорный элемент на его финальное место.
```go
    arr[left], arr[right] = arr[right], arr[left]
```
Рекурсивно вызываем quickSort для левой и правой частей.
```go
    quickSort(arr[:left])
    quickSort(arr[left+1:])
}
```

## Запуск программы

```bash
go run L1.16.go
```
**Вывод**
```bash
Неотсортированный массив: [1234 23 2342 233444 2 -123 0 34 234]
Отсортированный массив:   [-123 0 2 23 34 234 1234 2342 233444]
```
