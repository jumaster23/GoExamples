package main

import (
	"fmt" 
	"strings"
)

func main() {
	greet("Manuel", "Juma")
	name := "rosmery"
	toupperCase(&name)
	fmt.Println(name)	

}

func greet(nombre string, apellido string){
	fmt.Println("Hola", nombre, apellido)
}

func toupperCase(texto *string){
	*texto = strings.ToUpper(*texto)
	
}