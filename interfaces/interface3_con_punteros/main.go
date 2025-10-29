package main

import "fmt"

type storager interface {
	Get() string
	Set(string)
}

type Person struct {
	Name string
}

func NewPerson(name string) *Person {
	return &Person{name}
}

func (p *Person) Get() string {
	return p.Name
}

func (p *Person) Set(name string) {
	p.Name = name
}

func Exec(s storager, name string){
	s.Set(name)
	fmt.Println(s.Get())
}

func main() {
	p := NewPerson("Manuel")
	Exec(p, "Rosmery")
}