package payment

// ● Se o pagamento for para uma nova associação de membro, ative essa associação.
type ActivateMembershipSubscriber struct {
}

func (s *ActivateMembershipSubscriber) update(payment *payment) {
	payment.product.IsMember = true
}
