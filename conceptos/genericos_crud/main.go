package main

import (
	"fmt" 
	
)

type Productos[T uint | string] struct{

	ID T
	Descripcion string
	Price float64
}

func main() {
	producto1 := Productos[uint]{ID: 1, Descripcion:"Zapatos", Price: 200}
	fmt.Println(producto1)
}
