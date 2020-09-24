package main 

import "fmt"
import "os"
import "strconv"

func main() {
	//fmt.Println(os.Args)
	if len(os.Args) != 3 {
		fmt.Println("I need 2 integers through CLI")
		os.Exit(1)
	}

	a, _ := strconv.Atoi(os.Args[1])
	b, _ := strconv.Atoi(os.Args[2])

	c := a + b

	fmt.Println("sum =", c)
}
