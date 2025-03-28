package main

import (
	"errors"
	"fmt"
)

var ErrFull = errors.New("Queue is full")
var ErrEmpty = errors.New("Queue is empty")

// начало решения

// Queue - FIFO-очередь на n элементов
type Queue struct {
	data chan int
}

// Get возвращает очередной элемент.
// Если элементов нет и block = false -
// возвращает ошибку.
func (q Queue) Get(block bool) (int, error) {
	if block {
		return <-q.data, nil
	}

	select {
	case val := <-q.data:
		return val, nil
	default:
		return 0, ErrEmpty
	}

}

// Put помещает элемент в очередь.
// Если очередь заполнения и block = false -
// возвращает ошибку.
func (q Queue) Put(val int, block bool) error {
	if block {
		q.data <- val
		return nil
	}
	select {
	case q.data <- val:
		return nil
	default:
		return ErrFull
	}
}

// MakeQueue создает новую очередь
func MakeQueue(n int) Queue {
	return Queue{
		data: make(chan int, n), // Буферизированный канал размером n
	}
}

// конец решения

func main() {
	q := MakeQueue(2)

	err := q.Put(1, false)
	fmt.Println("put 1:", err)

	err = q.Put(2, false)
	fmt.Println("put 2:", err)

	err = q.Put(3, false)
	fmt.Println("put 3:", err)

	res, err := q.Get(false)
	fmt.Println("get:", res, err)

	res, err = q.Get(false)
	fmt.Println("get:", res, err)

	res, err = q.Get(false)
	fmt.Println("get:", res, err)
}
