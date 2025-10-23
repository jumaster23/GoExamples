package main

import "fmt"

var lista = []int{1,2,3,6,4}

func main() {
	division(100,10)
	division(200,25)
	division(34,0)
	division(200,25)
	fmt.Println(len(lista))
}

func division(dividenv, dividor int) {
	defer func ()  {
		if r := recover(); r != nil{
			fmt.Println("Me recupere del panico")
		}
	}()
	validateCero(dividor)
	fmt.Println(dividenv/dividor)
}


func validateCero(divisor int){
	if divisor == 0{
		panic("no se puede dividir entre cero")
	}

}