package main

import (
	"fmt"
	cursos "project/curso"
)

func main() {
	//Go := Curso{
	Go := cursos.New("Go desde cero", 12.34, false)
	Go.SetUserId([]uint{12, 56, 89}) 
	Go.Setclasses(map[uint]string{
			1: "introduccion",
			2: "estructuras",
			3: "maps",
		})
	
	Go.PrintClasses()
	Go.Setname("Poo con Go")
	fmt.Println(Go.Name())
	
}
