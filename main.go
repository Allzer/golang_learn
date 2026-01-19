package main

import (
	"fmt"
	"time"
)

func main() {
	
	ch := make(chan int)

	go func() {
		time.Sleep(2*time.Second)
		ch <- 42
	}()

	select {
	case <- ch:
		fmt.Println("Получено значение из канала")
	}
}