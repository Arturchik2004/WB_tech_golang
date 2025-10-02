package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func stubbornCat(wg *sync.WaitGroup) {
	defer func() {
		fmt.Println("Снова тихая ночь во дворе...")
		wg.Done()
	}()

	fmt.Println("Упрямый кот из первого способа вернулся на крышу: 'Я же говорил, что вернусь!'")
	fmt.Println("Кот: 'Никто меня отсюда не прогонит! Это МОЯ крыша!'")
	fmt.Println()

	for i := 1; i <= 3; i++ {
		fmt.Printf("Упрямый кот (попытка %d): МЯЯЯЯЯУ!!! \n", i)
		time.Sleep(500 * time.Millisecond)
	}

	fmt.Println()
	fmt.Println("Разъяренный сосед швыряет ботинок")
	fmt.Println("ЛЕТЯЩИЙ БОТИНОК ПОПАДАЕТ В ЦЕЛЬ!")

	runtime.Goexit() // Экстренно завершаем горутину ("Летящий ботинок судьбы")

	// Эти строки никогда не выполнятся - кот уже "вышел из чата"
	fmt.Println("Кот: 'Почему вокруг так тихо?'")
	fmt.Println("Кот: 'Мяу? Мяу?? Кто-нибудь меня слышит???'")
}

func main() {
	var wg sync.WaitGroup

	fmt.Println("Тихая ночь во дворе...")
	fmt.Println("Все соседи уже спят после ночного инцидента с котами из первого способа.")
	fmt.Println()

	wg.Add(1)
	go stubbornCat(&wg)

	wg.Wait()

}
