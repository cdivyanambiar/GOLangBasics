package concurrency

import (
	"fmt"
	"math/rand"
	"time"
)

func GetRandomNum() {
	nums := [] int {}
	rand.Seed(time.Now().Unix())
	 for i:=0 ; i<10; i++ {
	 	r := rand.Intn(99) + 1
	 	nums = append(nums, r)
	 }
	fmt.Println("Random number", nums)
}
