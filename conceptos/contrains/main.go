package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type myint int

func main() {
	var num1 myint = 2
	var num2 myint = 10
	fmt.Println(suma(num1,num2))
	fmt.Println(suma(50,62.5))

	fmt.Println(include([]int{1,2,3}, 3))
	fmt.Println(filter([]int{10,25,30},lesstoTwenty ))

}

func lesstoTwenty(num int)bool{

	return num < 20
}


type numbers interface{
	~int | ~float64 | ~float32 | ~uint
}

// 3 tipos de constrain tipo arbitraria, union elementos, de aproximacion de elementos
func suma[T constraints.Integer | constraints.Float](nums ...T) T{
	var total T

	for _, num := range nums{
		total += num
	}
	return total
}

func include[T comparable](list []T, value T)bool{
	for _, item := range list{
		if item == value{
			return true
		}
	}
	return false
}

func filter[T constraints.Integer | constraints.Float](nums []T, callback func(T) bool) []T {

	result := make([]T, 0, len(nums))

	for _, num := range nums {
		if callback(num) {
			result = append(result, num)
		}
	}

	return result
}
