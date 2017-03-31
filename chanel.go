package learn

import (
	//"time"

	"fmt"
)

var quit = make(chan bool)

func Producer(queue chan<- int) {

	for i := 0; i < 10000; i++ {
		fmt.Println("发送", i)
		queue <- i
	}
	close(queue)
}

func Consumer(queue <-chan int) {
	for i := range queue {
		fmt.Println("接收到", i)
	}
	quit <- true
}

func my() {
	queue := make(chan int)
	go Producer(queue)
	go Consumer(queue)
	<-quit

}
func myloop() {
	for i := 0; i < 10000; i++ {
		fmt.Println(i)
	}
}
