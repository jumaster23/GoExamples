package main

import "fmt"

//import "fmt"

func main() {
	character := "Juan"

	if character == "Manuel"{
		fmt.Println("Es Manuel")
	}else if character == "Juan"{
		fmt.Println("Es Juan")
	}else{
		fmt.Println("Es alguien mas")
	}


	jugador := "Lebron"

	switch jugador{
	case "MJ", "Lebron":
		fmt.Println("NBA")
	case "Vladimir", "Juan Soto":
		fmt.Println("MLB")
	default:
		fmt.Println("No es nadie")
	}
	
}