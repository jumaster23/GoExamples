package main

import "fmt"

type Person struct{
		Name string
		Age int8
		Haschilden bool
	}


func main() {
	Manuel := Person{
		Name: "Manuel",
		Age: 30,
		Haschilden: true,
	}

	//Mariluz := Person{"mariluz", 29, true}

	//fmt.Printf("%+v", Manuel)
	//fmt.Printf("%+v", Mariluz)

	Juan := &Manuel
	Juan.Age = 40
	fmt.Printf("%+v", Juan)
	
}
