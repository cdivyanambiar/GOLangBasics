package Inheritance

import "fmt"

type Person struct {
	Fname string
	Lname string
	Email string
}

func (p Person) String() string {
	return fmt.Sprintf("Person{%s,%s,%s}",p.Fname,p.Lname,p.Email)
}

func (p Person) Fullname() string {
	return p.Fname + " " + p.Lname
}

func (p Person) GetEmail() string {
	return p.Email
}

type Employee struct {
	P Person
	Eid int
	Salary float32
}

func (e Employee) String() string  {
	return fmt.Sprintf("Employee{%s,%s,%s, %d,%.1f}",e.P.Fname,e.P.Lname,e.P.Email,e.Eid,e.Salary)
}

func (e Employee) GetSalary() float32 {
	return  e.Salary
}
