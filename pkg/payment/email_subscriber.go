package payment

import "fmt"

// ● Se o pagamento for para uma adesão ou upgrade, envie um e-mail ao proprietário e
// informe-o sobre a ativação/upgrade.
type EmailSubscriber struct {
}

func (s *EmailSubscriber) update(payment *payment) {
	fmt.Println(fmt.Printf("send email when %s", payment.product.Category))
}
