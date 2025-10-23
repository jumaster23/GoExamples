package main

import (
	"fmt" 
	
)

func main() {
	fmt.Println(sum(2))
	fmt.Println(sum(2,3))
	fmt.Println(sum(4,2,1))
	fmt.Println(sum(10,5,7,4,6))

	greet := func (a string){ fmt.Println("Hola", a)}

	greet("Manuel")

}

func sum(nums ...int) (total int){

	for _, num := range nums{
		total += num
	}
	return 
}


