package main

import (
	"fmt"
	"sync"
)

type SafeMap struct {
	mu   sync.Mutex
	data map[int]int
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		data: make(map[int]int),
	}
}

func (sm *SafeMap) Set(key, value int) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.data[key] = value
}

func (sm *SafeMap) Get(key int) (int, bool) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	val, ok := sm.data[key]
	return val, ok
}

func main() {
	sMap := NewSafeMap()
	var wg sync.WaitGroup
	numGoroutines := 100

	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func(key int) {
			defer wg.Done()
			sMap.Set(key, key*10)
		}(i)
	}

	wg.Wait()

	fmt.Println("Проверка записи:")
	for i := 0; i < 5; i++ {
		if val, ok := sMap.Get(i); ok {
			fmt.Printf("Ключ: %d, Значение: %d\n", i, val)
		}
	}

	fmt.Printf("\nВсего записей в map: %d\n", len(sMap.data))
}
