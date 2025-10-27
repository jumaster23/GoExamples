package main

import "fmt"

type Person struct {
	Name string
	Age  uint
}

func NewPerson(name string, age uint) Person {
	return Person{name, age}
}

func (p Person) Greet() {
	fmt.Println("Hello")
}

type Employee struct{
	Person
	salary float64
}

func NewEmployee(name string, age uint, salary float64) Employee{
	return Employee{NewPerson(name,age), salary}
}

func (e Employee) Payroll(){
	fmt.Println(e.salary * 0.90)
}


func main()  {
	e := NewEmployee("Alejandro", 30, 100000)
	e.Payroll()
}
