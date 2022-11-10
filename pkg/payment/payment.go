package payment

import (
	"fmt"

	"github.com/Jonss/bhub-payment/pkg/product"
	"github.com/Jonss/bhub-payment/pkg/shipment"
)

type Payment struct {
	observers []Observer
	amount    uint64
	products  []product.Product
	shipment  *shipment.Shipment
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

	if p.containsProduct(learningToSki, product.VideoCategory) {
		subs := &SkiLearningSubscriber{}
		p.register(subs)
	}

	p.notifyAll()
}

func (p *Payment) containsProduct(productName string, category product.Category) bool {
	for _, product := range p.products {
		if product.Name == productName && product.Category == category {
			return true
		}
	}
	return false
}
