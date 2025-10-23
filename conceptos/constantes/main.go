package main

import "fmt"

const (
	computadora = "legion 5 pro"
	tarjeta = "RTX 3070ti"
)

const (
	jan = 1 // jan = iota + 1 lo que hara es comenzar en 1 y sumar 1 a cada constante
	feb = 2
	mar = 3
	apr = 4
	may = 5
	jun = 6

)


func main() {
	fmt.Println(computadora, tarjeta)
	fmt.Println(jan, feb, mar, apr, may, jun)
}