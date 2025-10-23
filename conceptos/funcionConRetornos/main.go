package main

import (
	"fmt" 
	"strings"
)

func main() {
	resultado := suma(50, 20)
	fmt.Println(resultado)

	lower, upper := convert("MaNuel")
	fmt.Println(lower, upper)

}

func suma(a, b int) int {

	return a + b
}

//func convert(text string)(lower string,upper string){ es recomendable cuando las funciones son pequenas
func convert(text string)(string, string){
	lower := strings.ToLower(text)
	upper := strings.ToUpper(text)

	return lower, upper
}
