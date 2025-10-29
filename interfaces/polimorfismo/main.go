package main

import "fmt"

type Paymethod interface {
	Pay()
}

//Implementando pago con paypal
type Paypal struct {
}

func (p Paypal) Pay() {
	fmt.Println("Pagado por Paypal")
}

//implementando pago Cash

type Cash struct{}

func (c Cash) Pay() {
	fmt.Println("Pagado con efectivo")
}

//implementando tarejta de credito

type CreditCard struct{}

func (cr CreditCard) Pay() {
	fmt.Println("Pagado con Credit Card")
}

//implementando la fabrica
func Factory(method uint) Paymethod{
	switch method{
	case 1:
		return Paypal{}
	case 2:
		return Cash{}
	case 3:
		return CreditCard{}
	default:
		return nil
	}
}



func main() {
	var input uint
	fmt.Println("Digita tu metodo de pago ")
	fmt.Println("1.Paypal ")
	fmt.Println("2. Cash ")
	fmt.Println("3. CreditCard")
	_, err := fmt.Scanln(&input)
	if err != nil{
		panic("Debe ingresar un metodo valido")
	}
	
	if input > 3{
		panic("Debe ingresar un metodo valido")
	}

	Paymethod := Factory(input)
	Paymethod.Pay()
}