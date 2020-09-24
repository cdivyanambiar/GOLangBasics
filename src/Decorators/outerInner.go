package Decorators

import "fmt"
/*
The private methood inner will return an inner function.
 */

func outer() func()int {
	state := 10
	inner := func() int {
		return state
	}
	return  inner
}

func OuterInner(){
	inn := outer()
	s := inn()
	fmt.Println(s)
}