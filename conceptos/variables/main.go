package main

import "fmt"

func main() {
	//var nombre string = "Milen"
	//var cedula int = 12345678
	//var genero string = "Masculino"

	var (
		nombre string = "Milen"
		cedula int = 12345678
		genero string = "Masculino"
	)

	var pais, provincia, ciudad = "Republica Dominicana", "Distrito Nacional", "Santo Domingo"
	empresa, rnc, direccion := "Empresa", 12345678, "Direccion"
	
	fmt.Println(nombre, cedula, genero)
	fmt.Println(pais, provincia, ciudad)
	fmt.Println(empresa, rnc, direccion)
	
}