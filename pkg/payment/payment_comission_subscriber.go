package payment

import "fmt"

// ● Se o pagamento for para um produto físico ou um livro, gere um pagamento de
// comissão ao agente.

type PaymentComissionSubscriber struct{}

func (s *PaymentComissionSubscriber) update(payment *Payment) {
	fmt.Println("PaymentComissionSubscriber")
}
