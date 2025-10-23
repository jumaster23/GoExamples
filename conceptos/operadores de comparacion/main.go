package main

import "fmt"

func main() {
	// operaciones comparacion : ==, !=, >, <, >=, <=
	fmt.Println((4 + 5) > 6)

	// operaciones logicas : &&, ||, !
	var age int = 33
	fmt.Println("Es adulto>=?", age >= 18 && age <= 60)

	//operadores unario !
	fmt.Println("Es igual?", !(4 == 4))

}
