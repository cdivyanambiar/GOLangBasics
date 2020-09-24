package concurrency

import (
	"fmt"
	"time"
)

func add(x, y int, ch chan int) {
	s := x + y
	fmt.Println("Computing ......")
	time.Sleep(5 * time.Second)
	ch <- s // Writing to channel ch
}
func ChannelDemo() {
	var a,b int
	c := make(chan int)
	a = 10
	b = 20
	go add(a, b, c)
	fmt.Println("Main waiting for sum.... ")
	r := <-c
	fmt.Println("Sum =",r)
}