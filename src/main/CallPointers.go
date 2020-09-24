package main

import (
	"Pointers"
	"fmt"
)

func main() {
	var m int = 10
	var p  * int
	p = &m
	*p = 20
	fmt.Println("m=",m,"*p =",*p)

	p = new(int)
	*p = 30
	fmt.Println("m=",m,"*p =",*p)

	b := [] int {10,20,30}
	Pointers.WithoutpointerEx(b)
	fmt.Println("in main b =",b)
	Pointers.PointerEx(&b)
	fmt.Println("in main b =",b)
}
