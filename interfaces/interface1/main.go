package main

import "fmt"

type Greeter interface {
	Greet()
}

type Person struct {
	Name string
}

func (p Person) Greet() {
	fmt.Printf("Hola soy %s", p.Name)
}

type Text string

func (t Text) Greet() {
	fmt.Printf("Hola soy %s", t)
}

func (t Text) Bye() {
	fmt.Printf("Adios soy %s", t)
}


type GreeterByer interface{
	Greeter
	Byer
}

type Byer interface{
	Bye()
}
func (p Person) Bye() {
	fmt.Printf("Adios %s", p.Name)
}

func All(gbs ...GreeterByer){
	for _, gs := range gbs{
		gs.Greet()
		gs.Bye()
	}
}

func main() {
	//var g Greeter = Person{"Alejandro"}
	//var g Greeter = Text("Manuel")
	p := Person{Name: "Alejandro"}
	var t Text = "Daisy"

	All(p, t)
	
	
}