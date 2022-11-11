package main

import (
	"github.com/Jonss/bhub-payment/pkg/payment"
	"github.com/Jonss/bhub-payment/pkg/product"
)

func main() {
	p1 := payment.NewPayment(1000, product.Product{
		Name:     "Os Lusíadas",
		Category: product.BookCategory,
	})
	p1.Execute()

}
