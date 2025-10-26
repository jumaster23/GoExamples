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

type Human struct{
	Age uint
	Children uint
}

func NewHuman(age, children uint)Human{
	return Human{age, children}
}

type Employee struct{
	Person
	Human
	salary float64
}

func NewEmployee(name string, age, children uint, salary float64) Employee{
	return Employee{NewPerson(name,age), NewHuman(age, children),salary}
}

func (e Employee) Payroll(){
	fmt.Println(e.salary * 0.90)
}


func main()  {
	e := NewEmployee("Alejandro", 30, 2, 100000)
	e.Payroll()
	e.Person.Greet()
	fmt.Println(e.Name, e.Human.Age)
}
