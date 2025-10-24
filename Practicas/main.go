package main

import (
	"fmt"
	"strings"
	
)

func main() {
	

	res := convert("Rio San Juan")


	for _, v := range res{
		fmt.Println(v)
	}
	result := []string{"Manuel","Juan","Calvin", "Fermin"}
	 result = Parpun(result)

	 fmt.Println(result)
}
	
	





func convert(palabra string) []string{
	
	lower := strings.ToLower(palabra)
	upper := strings.ToUpper(palabra)

	resultado := []string{lower, upper}

	
	return resultado
	
}



