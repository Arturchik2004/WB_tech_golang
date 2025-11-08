package main

import (
	"fmt"
	"reflect"
)

func identifyType(v interface{}) {
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
			fmt.Printf("Передано значение: %v, Тип: chan (канал), конкретный тип: %s\n", v, t.String())
		} else {
			fmt.Printf("Передано значение: %v, Тип: неизвестный (%s)\n", v, t.String())
		}
	}
}

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
