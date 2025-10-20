package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var st string

	fmt.Println("Ввести с клавиатуры - Y")
	fmt.Println("Взять с примера -25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5 - N")

	fmt.Scanln(&st)

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

	groups := make(map[int][]float64)

	for _, t := range temperatures {

		key := int(t/10) * 10

		groups[key] = append(groups[key], t)
	}

	fmt.Println(groups)
}
