package main

import (
	"fmt"
	"time"
)

func channel() {
	ch := make(chan string)
	go func() {
		time.Sleep(time.Second)
		// 将数据放入channel
		ch <- "hello , msg from channel"
	}()

	// 将channel 中的数据放入 msg
	msg := <-ch

	fmt.Println(msg)
}

func Select() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(time.Second)
		ch1 <- "msg from ch1"
	}()

	go func() {
		time.Sleep(time.Second)
		ch2 <- "msg from ch2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg, ok := (<-ch1):
			if ok {
				fmt.Println(msg)
			} else {
				fmt.Println("nil")
			}
		case msg := <-ch2:
			fmt.Println(msg)
		}
	}
}

func main() {
	// channel()
	Select()
}
