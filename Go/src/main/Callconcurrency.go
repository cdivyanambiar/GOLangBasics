package main

import "concurrency"

func main() {
	/*
	Parellel
	 */
	//concurrency.SayHi()
	//concurrency.SayHello()
	/*
	With Concurrency
	 */
	//go concurrency.SayHi()
	//go concurrency.SayHello()
	//for k:=0;k<10; k++ {
	//	fmt.Println("From main", k)
	//	time.Sleep(time.Second)
	//}

	/*
	calling channel
	 */
	// concurrency.ChannelDemo()

	//concurrency.AddNImpl()
	//concurrency.GetRandomNum()

	// concurrency.RandGame()

	link := make(chan string)
	done := make(chan bool)
	go concurrency.Producer(link)
	go concurrency.Consumer(link, done)
	<-done

}
