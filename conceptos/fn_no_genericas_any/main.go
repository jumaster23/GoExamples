package main

import "fmt"

func main() {
	printlist("edtem", "go", "rio san juan")
	printlist(1,2,3,5,4,8)

	fmt.Println(suma(2,5,7))
}

func printlist(list ...any) {
	for _, item := range list {
		fmt.Println(item)
	}
}

func suma(nums ...int) int {
	var total int

	for _, num := range nums{
		total += num
	}
	return total
}
