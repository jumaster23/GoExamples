package main

import (
	"fmt"
	cursos "project/curso"
)

func main() {
	//Go := Curso{
	Go := cursos.New("Go desde cero", 12.34, false)
	Go.UserID = []uint{12, 56, 89}
	Go.Classes = map[uint]string{
			1: "introduccion",
			2: "estructuras",
			3: "maps",
		}
	
	fmt.Printf("%+v\n",Go)
}
