package payment

import "fmt"

// Se o pagamento for para um produto físico, gerar uma guia de remessa para envio.
type ShipmentSubscriber struct {
}

func (o *ShipmentSubscriber) update(p *Payment) {
	fmt.Println("ShipmentSubscriber")
}
