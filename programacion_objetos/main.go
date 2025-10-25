package main

import (
	"fmt"
	cursos "project/curso"
)

func main() {
	//Go := Curso{
	Go := &cursos.Curso{
		Name:   "Go desde cero",
		Precio: 12.40,
		IsFree: false,
		UserID: []uint{12, 56, 89},
		Classes: map[uint]string{
			1: "introduccion",
			2: "estructuras",
			3: "maps",
		},
	}
	//PrintClasses(Go)
	Go.PrintClasses()
	//(&Go).Changeprices(67.12)
	Go.Changeprices(67.12)
	fmt.Println(Go.Precio)

	//CSS := Curso{Name:"CSS desde cero", Precio:40.5}
	//Javascript := Curso{}
	//Javascript.Name = "Curso JS"
	//Javascript.UserID = []uint{12,50,60}
}
