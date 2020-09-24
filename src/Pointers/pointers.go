package Pointers

import "fmt"
/*
Method in go is pass by reference only by pointers.

 */

func WithoutpointerEx(a [] int) {
	a = append(a, 100)
	fmt.Println("Inside f, a =", a)
}

func PointerEx(a * [] int) {
	*a = append(*a, 100)
	fmt.Println("Inside f, a =", *a)
}

