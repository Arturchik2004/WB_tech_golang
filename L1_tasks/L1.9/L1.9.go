package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func res(elem <-chan int, result chan<- int) {
	for val := range elem {
		result <- val
	}
	close(result)
}
func res2(results <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range results {
		fmt.Println(val * val)
	}
}

func main() {
	elem := make(chan int)
	result := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go res2(result, &wg)
	go res(elem, result)

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

	for _, val := range m {
		elem <- val
	}

	close(elem)

	wg.Wait()

}
