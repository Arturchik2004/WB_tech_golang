# L1.10 Группировка температур на Go
<img width="803" height="506" alt="image" src="https://github.com/user-attachments/assets/cf73504c-56c3-4ca6-a4b8-7e265fabb6e8" />


Этот код демонстрирует группировку температурных значений в диапазоны с шагом 10 градусов.

## Описание

Программа выполняет следующие действия:
1.  Предлагает пользователю выбрать источник данных: ввести температуры вручную или использовать предопределенный набор.

    ```go
    var st string

    fmt.Println("Ввести с клавиатуры - Y")
    fmt.Println("Взять с примера -25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5 - N")

    fmt.Scanln(&st)
    ```

2.  В зависимости от выбора, либо считывает данные с клавиатуры, либо использует встроенный срез.

    ```go
    var temperatures []float64

    if st == "Y" {
        fmt.Println("Введите температуры через пробел:")
        reader := bufio.NewReader(os.Stdin)
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)
        var temp float64
        var err error
        for _, s := range strings.Split(input, " ") {
            if s != "" {
                if temp, err = strconv.ParseFloat(s, 64); err == nil {
                    temperatures = append(temperatures, temp)
                }
            }
        }
    } else {
        temperatures = []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
    }
    ```

3.  Группирует температуры в `map`, где ключ — это начало диапазона (например, -20, 10), а значение — срез температур в этом диапазоне.

    ```go
    groups := make(map[int][]float64)

    for _, t := range temperatures {

        key := int(t/10) * 10

        groups[key] = append(groups[key], t)
    }
    ```

4.  Выводит результат в консоль.

    ```go
    fmt.Println(groups)
    ```

## Как запустить

1.  Убедитесь, что у вас установлен Go.
2.  Перейдите в директорию с файлом:
    ```bash
    cd L1_tasks/L1.10
    ```
3.  Запустите программу:
    ```bash
    go run L1.10.go
    ```
4.  Следуйте инструкциям в консоли.

## Пример работы

**Ввод (выбираем 'N'):**
```
N
```

**Вывод:**
```
map[-20:[-25.4 -27 -21] 10:[13 19 15.5] 20:[24.5] 30:[32.5]]
```
