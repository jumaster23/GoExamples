package invoice

import (
	"project/pkg/customer"
	"project/pkg/invoiceitem"
)

type Invoice struct {
	country string
	ciudad  string
	total   float64
	client  customer.Customer
	items   []invoiceitem.Item
}

// return a new invoice
func New(country, ciudad string, client customer.Customer, items []invoiceitem.Item)Invoice{
	return Invoice{
		country: country,
		ciudad: ciudad,
		client: client,
		items: items,
	}
}

// settotal is the setter of Invoice.total
func (i *Invoice) Settotal(){
	for _, item := range i.items{
		i.total += item.Value()
	}
}