package main

import "fmt"

type curso struct {
	name string
}

func (c curso) Print() {
	fmt.Printf("%+v\n", c)
}

//declaracion de alias
//type myAlias = curso

//deficion de tipo

type newCurso curso

type newBool bool

func (b newBool) string() string{
	if b{
		return "VERDADERO"
	}
	return "FALSO"
}


func main()  {
	//c := newCurso{name: "Manuel"}
	//fmt.Printf("El tipo es: %T\n", c)
	var b newBool = true
	fmt.Println(b.string())	
}

