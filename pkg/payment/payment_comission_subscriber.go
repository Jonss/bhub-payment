package payment

import "fmt"

// ● Se o pagamento for para um produto físico ou um livro, gere um pagamento de
// comissão ao agente.

type PaymentComissionSubscriber struct{}

func (s *PaymentComissionSubscriber) update(payment *payment) {
	fmt.Println("chega aqui")
	commissionAmount := int64(float64(payment.amount) * 0.10)
	payment.comission = &commissionAmount
}
