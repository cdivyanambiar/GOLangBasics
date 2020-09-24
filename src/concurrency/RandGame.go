package concurrency

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type HoldData struct {
   data [] int
}

func (h HoldData) CalculateFrequency() {
	keys := make(map[int]int)
	for _, entry := range h.data {
		keys[entry]++
	}
	fmt.Println(keys)
}

func IsPrime(n int) bool {
	if n < 2 { return  false }
	i :=2
	for i<n {
		if n%i == 0 {
			return true
		}
		i = i + 1
	}
	return false
}

func GetRandPrimes(id int, count int, c chan int, prime chan [] int)   {
	fmt.Println("goRoutine...", id, "started")
	out := make([] int, 0)
    primes := [] int{}
    i := 0
    for i<count {
    	min := 10
		max := 999
		r := rand.Intn(max - min) + min
   	    out = append(out, r)
   	    if IsPrime(r) {
   	 		primes = append(primes, r)
   	 	    i = i  + 1
	 }
   }
   time.Sleep(5 * time.Second)
    c<-id
	prime <- out
}

func RandGame() {
	if len(os.Args) != 2 {
		fmt.Println("I need an integer through cli")
	}
	rand.Seed(time.Now().Unix())
	n, _ := strconv.Atoi(os.Args[1])
	c := make(chan int)
	out := make(chan [] int)


	d := make([]int,0)
	fmt.Println(d)
	go GetRandPrimes(1, n,c,out)
	go GetRandPrimes(2, n,c,out)
	go GetRandPrimes(3, n,c,out)

	data_obj := HoldData {d}

	for i := 1; i < 4; i++ {
		select {
		case r1 := <-c:
			primeOut1 := <-out
			for _,dt := range primeOut1 {
				d = append(d, dt)
			}
			fmt.Println("Go routine ....", r1, "Finished -------> ", i)

		case r2 := <-c:
			primeOut2 := <-out
			for _,dt := range primeOut2 {
				d = append(d, dt)
			}
			fmt.Println("Go routine ....", r2, "Finished -------> ", i)

		case r3 := <-c:
			primeOut3 := <-out
			for _,dt := range primeOut3 {
				d = append(d, dt)
			}
			fmt.Println("Go routine ....", r3, "Finished -------> ", i)

		case t := <-time.After(5 * time.Second):
			fmt.Println("Time", t)

		}
	}

	data_obj.data = d
	data_obj.CalculateFrequency()
}