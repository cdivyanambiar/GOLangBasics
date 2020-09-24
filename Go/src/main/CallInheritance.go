package main

import (
	"Inheritance"
	"fmt"
)

func main() {
	p := Inheritance.Person{"Divya","Nambiar","dnambiar@gmail.com"}
	fmt.Println(p.GetEmail())

	emp := Inheritance.Employee{p,123,10000}
	fmt.Println(emp)

	fn := emp.P.Fullname()
	fmt.Println(fn)

	sal := emp.GetSalary()
	fmt.Println(sal)
}

