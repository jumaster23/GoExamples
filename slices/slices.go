package slices

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

const pkgName = "slices"


func Log(msg string){
	fmt.Println("LOG: ", msg)
}

func Filter[T constraints.Ordered](nums []T, callback func(T) bool) []T{

	result := make([]T, 0, len(nums))

	for _, num := range nums {
		if callback(num) {
			result = append(result, num)
		}
	}

	return result
}