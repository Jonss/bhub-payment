package payment

import (
	"fmt"

	"github.com/Jonss/bhub-payment/pkg/product"
)

type payment struct {
	observers   map[string]Observer
	amount      int64
	product     product.Product
	packingSlip *PackingSlip
	comission   *int64
}

func NewPayment(amount int64, product product.Product) payment {
	return payment{
		observers: make(map[string]Observer),
		product:   product,
		amount:    amount,
	}
}

type PackingSlip struct {
	title string
}

func (p *payment) notifyAll() {
	for _, o := range p.observers {
		o.update(p)
	}
}

func (p *payment) register(name string, o Observer) {
	p.observers[name] = o
}

func (p *payment) Execute() {
	fmt.Println("executes payment")

	//  ● Se o pagamento for para o vídeo específico “Aprendendo a Esquiar”, adicione um
	// vídeo gratuito de “Primeiros Socorros” à guia de remessa (resultado de uma decisão
	// judicial em 1997).
	if p.product.Name == learningToSki && p.product.Category == product.VideoCategory {
		subs := &SkiLearningSubscriber{}
		p.register("ski_learning_subscriber", subs)
	}

	//    ● Se o pagamento for para um produto físico, gerar uma guia de remessa para envio.
	if p.product.Category == product.PhysicalMediaCategory {
		subs := &ShipmentSubscriber{}
		p.register("shipment_subscriber", subs)
	}

	// ● Se o pagamento for para um livro, crie uma guia de remessa duplicada para o
	// departamento de royalties.
	if p.product.Category == product.BookCategory {
		subs := &DoubleShipmentSubscriber{}
		p.register("double_packing_slip", subs)
	}

	// ● Se o pagamento for para uma nova associação de membro, ative essa associação.
	if p.product.Category == product.MemberAssociationCategory {
		subs := &ActivateMembershipSubscriber{}
		p.register("activate_membership_subscriber", subs)
	}

	// ● Se o pagamento for um upgrade para uma associação, aplique o upgrade.
	if p.product.Category == product.MemberAssociationUpgradeCategory {
		subs := &UpgradeMembershipSubscriber{}
		p.register("upgrade_membership_subscriber", subs)
	}
	// ● Se o pagamento for para uma adesão ou upgrade, envie um e-mail ao proprietário e
	// informe-o sobre a ativação/upgrade.
	if p.product.Category == product.MemberAssociationCategory ||
		p.product.Category == product.MemberAssociationUpgradeCategory {
		subs := &EmailSubscriber{}
		p.register("send_email_subscriber", subs)
	}

	// ● Se o pagamento for para um produto físico ou um livro, gere um pagamento de
	// comissão ao agente.
	if p.product.Category == product.PhysicalMediaCategory ||
		p.product.Category == product.BookCategory {
		subs := &PaymentComissionSubscriber{}
		p.register("payment_commission_subscriber", subs)
	}

	p.notifyAll()
}
