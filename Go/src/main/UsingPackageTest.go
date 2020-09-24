package main

import (
	"fmt"
	"test"
)

func main(){
	a := test.Add(2,3)
	fmt.Println(a)

	m := test.Multiply(4,5)
	fmt.Println(m)

	d := test.DefaultaddResulr()
	fmt.Println(d)
}