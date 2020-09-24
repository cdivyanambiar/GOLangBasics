package test

import "insidetest"

func Add(a,b int) int {
	return a + b
}

func DefaultaddResulr() int {
	x, y := insidetest.ReturnNum()
	return x + y
}



