package payment

import "fmt"

// ● Se o pagamento for para uma nova associação de membro, ative essa associação.
type UpgradeMembershipSubscriber struct {
}

func (s *UpgradeMembershipSubscriber) update(payment *Payment) {
	fmt.Println("UpgradeMembershipSubscriber")
}
