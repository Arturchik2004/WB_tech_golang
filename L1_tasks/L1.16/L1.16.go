package main

import (
	"fmt"
	"math/rand"
	"time"
)

func quickSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	left, right := 0, len(arr)-1
	rand.Seed(time.Now().UnixNano())
	pivotIndex := rand.Intn(len(arr))
	pivot := arr[pivotIndex]
	arr[pivotIndex], arr[right] = arr[right], arr[pivotIndex]
	for i := range arr {
		if arr[i] < pivot {
			arr[left], arr[i] = arr[i], arr[left]
			left++
		}
	}

	arr[left], arr[right] = arr[right], arr[left]

	quickSort(arr[:left])
	quickSort(arr[left+1:])
}

func main() {
	arr := []int{1234, 23, 2342, 233444, 2, -123, 0, 34, 234}

	fmt.Println("Неотсортированный массив:", arr)

	quickSort(arr)

	fmt.Println("Отсортированный массив:  ", arr)
}
