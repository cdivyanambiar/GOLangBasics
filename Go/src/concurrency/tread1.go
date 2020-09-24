package concurrency

import (
	"fmt"
	"time"
)

func SayHi() {
	for i:=0; i<10; i++ {
		fmt.Println("Say Hi", i)
		time.Sleep(time.Second)
	}
}

func SayHello() {
	for i:=0; i<10; i++ {
		fmt.Println("Say Hello", i)
		time.Sleep(time.Second)
	}
}
