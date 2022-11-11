package payment

// Se o pagamento for para um produto físico, gerar uma guia de remessa para envio.
type ShipmentSubscriber struct {
}

func (o *ShipmentSubscriber) update(p *payment) {
	p.packingSlip = &PackingSlip{title: "guia de remessa de envio"}
}
