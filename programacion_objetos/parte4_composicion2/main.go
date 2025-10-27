package main

import (
	"fmt"
	"project/pkg/customer"
	"project/pkg/invoice"
	"project/pkg/invoiceitem"
)



func main() {
	i := invoice.New(
		"Colombia", 
		"Popayan",
		customer.New("alejandro", "calle 123", "46464"),
		invoiceitem.Newitems(
			invoiceitem.New(1, "curso de go", 12.34),
			invoiceitem.New(2, "curso de Poo con go", 54.23),
			invoiceitem.New(3, "curso de testing con Go", 90.00),
		),
	)

	i.Settotal()
	fmt.Printf("%+v", i)
}
