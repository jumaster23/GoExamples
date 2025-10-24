package main

import (
	
	"strings"
)

func Parpun(palabras []string)[]string {
	for i := range palabras{
		palabras[i] = strings.ToUpper(palabras[i])

	}
	return palabras
}
