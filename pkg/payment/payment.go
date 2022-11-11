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

	// add events
	subscribeEvents(p)

	p.notifyAll()
}
