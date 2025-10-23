package main

import "fmt"

func main() {
	nums := []int{2, 12, 23, 98, 21, 79}

	result := filter(nums,lesstoTwenty )
	fmt.Println(result)

	suma := sum(2)(3)
	fmt.Println(suma)
}


func greatertoFifty (num int) bool { 
	return num > 50 

}

func lesstoTwenty(num int)bool{

	return num < 20
}

func filter(nums []int, callback func(int) bool) []int {

	result := make([]int, 0, len(nums))

	for _, num := range nums {
		if callback(num) {
			result = append(result, num)
		}
	}

	return result
}

func sum(a int) func(int)int{
	return func(b int) int{
		return a + b
	}
}
