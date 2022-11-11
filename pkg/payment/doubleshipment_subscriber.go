package payment

// ● Se o pagamento for para um livro, crie uma guia de remessa duplicada para o
// departamento de royalties.
type DoubleShipmentSubscriber struct {
}

func (s *DoubleShipmentSubscriber) update(payment *payment) {
	payment.packingSlip = &PackingSlip{title: "guia para o departamento de royalties"}
}
