package main

import (
	"fmt"
	"sort"
)

func heapyfy(arr []int, n int, i int) {
	l := i
	left := 2*i + 1
	right := 2*i + 2
	if left < n && arr[l] < arr[left] {
		l = left
	}
	if right < n && arr[l] < arr[right] {
		l = right
	}
	if l != i {
		arr[i], arr[l] = arr[l], arr[i]
		heapyfy(arr, n, l)
	}
}
func heap_sort(arr []int) {
	n := len(arr)
	for i := n/2 - 1; i >= 0; i-- {
		heapyfy(arr, n, i)
	}
	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapyfy(arr, i, 0)
	}
}

func main() {
	arr1 := []int{1234, 23, 2342, 233444, 2, -123, 0, 34, 234}
	arr := []int{1234, 23, 2342, 233444, 2, -123, 0, 34, 234}
	heap_sort(arr1)
	sort.Ints(arr)
	fmt.Println(arr1)
	fmt.Println(arr)
}
