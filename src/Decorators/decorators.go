package Decorators
/*
Decorators essentially allow you to wrap existing functionality
and append or prepend your own custom functionality on top
 */
import "fmt"

/*
Method with small letter starting in private
 */
func add10(f func(int, int) int) func (int,int) int {
	decor := func(x,y int) int {
		return f(x,y) + 10
	}
	return decor
}
func addInner(a, b int) int {
	return a + b
}

func DecoratorsDemo(){
	s := addInner(2,3)
	fmt.Println("Sum =",s)

	action := add10(addInner)
	s10 := action(2, 3)
	//OR
	s10 = add10(addInner)(2,3)
	fmt.Println("Sum =",s10)
}
