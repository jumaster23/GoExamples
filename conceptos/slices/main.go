package main

import "fmt"

func main() {
	things := [7]string{"galleta", "cafe", "pan", "leche", "queso", "carne", "pescado"}
	slcs := things[:]
	slcs[2] = "queso"

	//fmt.Println(things)
	//fmt.Println(slcs)


	animals := []string{"gato", "perro", "pato", "gallina", "vaca", "caballo", "oveja"}
	pets := animals[0:3]
	pets = append(pets, "gallina")

	//fmt.Println(pets)
	//fmt.Println(animals)
	//fmt.Println("Tamaños de los animales:",len(pets))
	//fmt.Println("Capacidad de los animales:",cap(pets))


	paises := make([]string, 0, 3)
	paises = append(paises, "Argentina", "Brasil", "Chile")
	paises = append(paises, "RD")
	fmt.Println(paises)
	fmt.Println("Tamaños de los paises:",len(paises))
	fmt.Println("Capacidad de los paises:",cap(paises))




}
// Los slices son apuntadores de array, no poseen datos