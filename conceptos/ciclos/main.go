package main

import "fmt"

var count = 10
var k = 1

func main()  {
	for i := 1; i <= count; i ++{
		fmt.Println(i)
	}

	j := 1

	for j <= count{
		fmt.Println(j)
		j++
	}

	for{
		if k == 10{
			break
		}
		fmt.Println("Valor de k", k)
		k++
	}

	food := []string  {"Pizza", "Pasta", "Filet Mignon"}

	for i, v := range food{
		fmt.Println("indice", i, "valor", v)
	} 

	numeros := []int8{2,4,6,7,8,9,4,5,12,}

	for _, v := range numeros{
		fmt.Println(v * 10)
	}

	numeros2 := []int8{2,4,6,7,8,9,4,5,12,}

	for i := range numeros2{
		numeros2[i] *= 2
	}
	fmt.Println(numeros2)

	food2 := map[string]string{
		"Nombre:": "Manuel",
		"Apellido:": "juma",
		"Direccion:": "Rio San Juan",
	}

	for key, v := range food2{
		fmt.Println(key, v)

	}
	
	
}