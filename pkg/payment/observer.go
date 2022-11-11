package payment

import "fmt"

type Observer interface {
	update(p *payment)
}

// paymentComissionSubscriber
type PaymentComissionSubscriber struct{}

func (s *PaymentComissionSubscriber) update(payment *payment) {
	commissionAmount := int64(float64(payment.amount) * 0.10)
	payment.comission = &commissionAmount
}

//end: paymentComissionSubscriber

// shipmentSubscriber
type ShipmentSubscriber struct {
}

func (o *ShipmentSubscriber) update(p *payment) {
	p.packingSlip = &PackingSlip{title: "guia de remessa de envio"}
}

//end: shipmentSubscriber

// skiLearningSubscriber
const learningToSki string = "Aprendendo a Esquiar"

type SkiLearningSubscriber struct {
}

func (s *SkiLearningSubscriber) update(payment *payment) {
	payment.packingSlip = &PackingSlip{title: "adicione um vídeo gratuito de “Primeiros Socorros”"}
}

//end: skiLearningSubscriber

// upgradeMembershipSubscriber
type UpgradeMembershipSubscriber struct {
}

func (s *UpgradeMembershipSubscriber) update(payment *payment) {
	payment.product.IsMember = true
}

//end: upgradeMembershipSubscriber

// activateMembershipSubscriber
type ActivateMembershipSubscriber struct {
}

func (s *ActivateMembershipSubscriber) update(payment *payment) {
	payment.product.IsMember = true
}

//end: activateMembershipSubscriber

// emailSubscriber
type EmailSubscriber struct {
}

func (s *EmailSubscriber) update(payment *payment) {
	fmt.Println(fmt.Printf("send email when %s", payment.product.Category))
}

//end: emailSubscriber

// doubleShipmentSubscriber
type DoubleShipmentSubscriber struct {
}

func (s *DoubleShipmentSubscriber) update(payment *payment) {
	payment.packingSlip = &PackingSlip{title: "guia para o departamento de royalties"}
}

//end: doubleShipmentSubscriber
