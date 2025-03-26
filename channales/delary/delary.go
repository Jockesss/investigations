package main

import (
	"fmt"
	"math/rand"
	"time"
)

// начало решения

func delay(dur time.Duration, fn func()) func() {
	cancel := make(chan struct{}, 1) // буферизированный канал
	done := make(chan struct{})

	go func() {
		select {
		case <-time.After(dur):
			fn()
		case <-cancel:
			// отмена — ничего не делаем
		}
		close(done)
	}()

	// возвращаем функцию отмены
	return func() {
		select {
		case cancel <- struct{}{}: // если канал пуст, мы можем отправить в него значение
		default:
			// если в канале уже есть значение, значит cancel() был вызван
		}
	}
}

// конец решения

func main() {
	work := func() {
		fmt.Println("work done")
	}

	cancel := delay(100*time.Millisecond, work)

	time.Sleep(10 * time.Millisecond)
	if rand.Float32() < 0.5 {
		cancel()
		fmt.Println("delayed function canceled")
	}
	time.Sleep(100 * time.Millisecond)
}
