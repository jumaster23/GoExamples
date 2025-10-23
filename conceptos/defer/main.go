package main

import (
	"os" 
	"fmt"
)


func main() {
	file, err := os.Create("test.txt")
	if err != nil{
		fmt.Println(err)
		return
		}
	defer file.Close()

	_, err = file.Write([]byte("hello Manuel Juma"))
	if err != nil{
		fmt.Println(err)
		return
		}

	
}
