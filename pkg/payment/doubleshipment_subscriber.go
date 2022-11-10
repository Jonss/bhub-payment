package payment

import "fmt"

// ● Se o pagamento for para um livro, crie uma guia de remessa duplicada para o
// departamento de royalties.
type DoubleShipmentSubscriber struct {
}

func (o *DoubleShipmentSubscriber) update(payment *Payment) {
	fmt.Println("DoubleShipmentSubscriber")
}
