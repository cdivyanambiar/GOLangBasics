package concurrency

import (
	"fmt"
	"time"
)

func  AddN(s []int, c chan int)  {
	result := 0
	for _,x := range s {
		result = result + x
		time.Sleep(time.Second)
	}
	c <- result
}
func AddNImpl() {
	data := []int {1, 2, 3, 4, 5, 6 ,7 ,8, 9, 10}
	n := len(data)

	c1 := make(chan int)
	c2 := make(chan int)

	go AddN(data[:n/2], c1)
    go AddN(data[n/2:], c2)

	r1 := <-c1
	r2 := <-c2

	fmt.Println("r1=",r1, "r2=", r2)

	for i:=0; i<2; i++ {
		select {
		case r1 := <-c1:
			fmt.Println("r1==", r1)
		case r2 := <-c2:
			fmt.Println("r2==", r2)

		case t := <-time.After(5 * time.Second):
			fmt.Println("Time", t)
		}
	}
}
