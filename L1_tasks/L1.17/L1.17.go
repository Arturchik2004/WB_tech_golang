package main

import (
	"fmt"
)

func binarySearch(arr []int, target int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := low + (high-low)/2

		guess := arr[mid]

		if guess == target {
			return mid
		}

		if guess > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}

func main() {
	sortedSlice := []int{1, 3, 5, 7, 9, 11, 13, 15}

	targets := []int{7, 1, 15, 4} // 4 нет в списке

	for _, t := range targets {
		index := binarySearch(sortedSlice, t)
		if index != -1 {
			fmt.Printf("Число %d найдено под индексом %d\n", t, index)
		} else {
			fmt.Printf("Число %d не найдено\n", t)
		}
	}
}
