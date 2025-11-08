# L1.14 Определение типа переменной в runtime

<img width="928" height="513" alt="image" src="https://github.com/user-attachments/assets/97bef6e5-7d94-43e0-aeea-9f9ad2324586" />


## Описание 

Для `int`, `string` и `bool` использую `switch v.(type)` напрямую. Для `chan` (канала) можно указать конкретные типы каналов (например, `chan int`), но чтобы определить любой тип канала используем `switch` для известных типов, а в блоке `default` применить пакет `reflect` для проверки "вида" переменной.

## Реализация
 `identifyType` принимает пустой интерфейс и определяет его конкретный тип
 оператор `switch v.(type)` нужен для определения типа
В `default` используем пакет `reflect`, чтобы проверить "вид" типа
 ```go
func identifyType(v interface{}) {
	// Используем 
	switch v.(type) {
	case int:
		fmt.Printf("Передано значение: %v, Тип: int\n", v)
	case string:
		fmt.Printf("Передано значение: %v, Тип: string\n", v)
	case bool:
		fmt.Printf("Передано значение: %v, Тип: bool\n", v)
default:
		t := reflect.TypeOf(v)
		if t.Kind() == reflect.Chan {
			// t.String() вернет полное имя типа, например "chan int" или "chan bool"
			fmt.Printf("Передано значение: %v, Тип: chan (канал), конкретный тип: %s\n", v, t.String())
		} else {
			fmt.Printf("Передано значение: %v, Тип: неизвестный (%s)\n", v, t.String())
		}
	}
}
```

***А дальше в main производим проверку функции***

```go
func main() {
	var a int = 100
	var b string = "Привет, Go!"
	var c bool = true
	var d chan int = make(chan int)
	var e chan string = make(chan string)
	var f float64 = 123.45

	fmt.Println("Тестирование функции")
	identifyType(a)
	identifyType(b)
	identifyType(c)
	identifyType(d)
	identifyType(e)
	identifyType(f)
}
```
## Запуск

```bash
go run L1.14.go
```

## Вывод
```
Тестирование функции
Передано значение: 100, Тип: int
Передано значение: Привет, Go!, Тип: string
Передано значение: true, Тип: bool
Передано значение: 0xc00001c0e0, Тип: chan (канал), конкретный тип: chan int
Передано значение: 0xc00001c150, Тип: chan (канал), конкретный тип: chan string
Передано значение: 123.45, Тип: неизвестный (float64)
```


