package main

import "fmt"

func main() {
	var color string = "rojo"


	var puntero *string = &color
	*puntero = "azul"


	fmt.Printf("tipo: %T, valor: %s, Direccion: %v\n", color, color, &color)
	fmt.Printf("tipo: %T, valor: %v, desreferenciacion: %s\n", puntero, puntero, *puntero)
}