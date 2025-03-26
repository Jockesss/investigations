// Ограничитель скорости
package main

import (
	"errors"
	"fmt"
	"time"
)

var ErrCanceled error = errors.New("canceled")

// начало решения

// throttle следит, чтобы функция fn выполнялась не более limit раз в секунду.
// Возвращает функции handle (выполняет fn с учетом лимита) и cancel (останавливает ограничитель).
// начало решения
func throttle(limit int, fn func()) (handle func() error, cancel func()) {
	// if limit <= 0 {
	// 	panic("limit must be > 0")
	// }

	// done := make(chan struct{})
	// interval := time.Second / time.Duration(limit)

	// go func() {
	// 	timer := time.NewTimer(interval)
	// 	defer close(done)
	// 	for {
	// 		select {
	// 		case <-done:
	// 			return
	// 		case <-timer.C:

	// 		}
	// 	}
	// }()

	// handle = func() error {

	// }

	// cancel = func() {

	// }

	return handle, cancel
}

// конец решения

func main() {
	work := func() {
		fmt.Print(".")
	}

	handle, cancel := throttle(5, work)
	defer cancel()

	start := time.Now()
	const n = 10
	for i := 0; i < n; i++ {
		handle()
	}
	fmt.Println()
	fmt.Printf("%d queries took %v\n", n, time.Since(start))
}
