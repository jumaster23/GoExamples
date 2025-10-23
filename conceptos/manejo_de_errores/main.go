package main

import (
	"errors"
	"fmt"
	"strconv"
)

var errNotFound = errors.New("not found")

var food = map[int]string{
	1: "Manzana",
	2: "Pera",
}

func main() {
	num, err := strconv.Atoi("10")
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(num)

	found, err := search("4")
	if errors.Is(err, errNotFound){
		fmt.Println("pudimos controlar el error correctamente")
		return
	}
	if err != nil{
		fmt.Println("Search()",err)
		return
	}
	fmt.Println(found)	
}

func search(key string)(string, error){
	num, err := strconv.Atoi(key)
	if err != nil{
		return "", fmt.Errorf("strconv.Atoi() %w", err)
	}
	a, err := findfood(num)
	if err != nil{
		return "", fmt.Errorf("findfood %w", err)
	}
	return a, nil
}

func findfood(id int)(string, error){
	value, ok := food[id]
	if !ok {
		return "", errNotFound
	}
	return value, nil
}

