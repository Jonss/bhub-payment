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

	p2 := payment.NewPayment(1000, product.Product{
		Name:     "Aprendendo a Esquiar",
		Category: product.VideoCategory,
	})
	p2.Execute()
}
