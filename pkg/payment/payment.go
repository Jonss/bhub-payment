package payment

import (
	"fmt"

	"github.com/Jonss/bhub-payment/pkg/product"
	"github.com/Jonss/bhub-payment/pkg/shipment"
)

type Payment struct {
	observers   []Observer
	amount      uint64
	product     product.Product
	shipment    *shipment.Shipment
	packingSlip *PackingSlip
}

type PackingSlip struct {
	title string
}

func (p *Payment) notifyAll() {
	for _, o := range p.observers {
		o.update(p)
	}
}

func (p *Payment) register(o Observer) {
	p.observers = append(p.observers, o)
}

func (p *Payment) Execute() {
	fmt.Println("executes payment")

	//  ● Se o pagamento for para o vídeo específico “Aprendendo a Esquiar”, adicione um
	// vídeo gratuito de “Primeiros Socorros” à guia de remessa (resultado de uma decisão
	// judicial em 1997).
	if p.product.Name == learningToSki && p.product.Category == product.VideoCategory {
		subs := &SkiLearningSubscriber{}
		p.register(subs)
	}

	//     ● Se o pagamento for para um produto físico, gerar uma guia de remessa para envio.
	if p.product.Category == product.PhysicalMediaCategory {
		subs := &ShipmentSubscriber{}
		p.register(subs)
	}

	// ● Se o pagamento for para um livro, crie uma guia de remessa duplicada para o
	// departamento de royalties.
	// ● Se o pagamento for para uma nova associação de membro, ative essa associação.
	// ● Se o pagamento for um upgrade para uma associação, aplique o upgrade.
	// ● Se o pagamento for para uma adesão ou upgrade, envie um e-mail ao proprietário e
	// informe-o sobre a ativação/upgrade.

	// ● Se o pagamento for para um produto físico ou um livro, gere um pagamento de
	// comissão ao agente.

	p.notifyAll()
}
