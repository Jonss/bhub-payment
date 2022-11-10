package payment

import "fmt"

// ● Se o pagamento for para uma nova associação de membro, ative essa associação.
type ActivateMembershipSubscriber struct {
}

func (s *ActivateMembershipSubscriber) update(payment *Payment) {
	fmt.Println("ActivateMembershipSubscriber")
}
