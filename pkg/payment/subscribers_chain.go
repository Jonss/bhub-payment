package payment

import (
	"github.com/Jonss/bhub-payment/pkg/product"
)

type Chain interface {
	execute(p *payment)
	setNext(c Chain)
}

// skiLearningChain
type SkiLearningChain struct {
	next Chain
}

func (c *SkiLearningChain) execute(p *payment) {
	if p.product.Name == learningToSki && p.product.Category == product.VideoCategory {
		subs := &SkiLearningSubscriber{}
		p.register("ski_learning_subscriber", subs)
	}
}

func (c *SkiLearningChain) setNext(next Chain) {
	c.next = next
}

// skiLearningChain
type PhysicalMediaChain struct {
	next Chain
}

func (c *PhysicalMediaChain) execute(p *payment) {
	if p.product.Category == product.PhysicalMediaCategory {
		subs := &ShipmentSubscriber{}
		p.register("shipment_subscriber", subs)
	}
	c.next.execute(p)
}

func (c *PhysicalMediaChain) setNext(next Chain) {
	c.next = next
}

// bookChain
type BookChain struct {
	next Chain
}

func (c *BookChain) execute(p *payment) {
	if p.product.Category == product.BookCategory {
		subs := &DoubleShipmentSubscriber{}
		p.register("double_packing_slip", subs)
	}
	c.next.execute(p)
}

func (c *BookChain) setNext(next Chain) {
	c.next = next
}

//end: bookChain

// memberAssociationChain
type MemberAssociationChain struct {
	next Chain
}

func (c *MemberAssociationChain) execute(p *payment) {
	if p.product.Category == product.MemberAssociationCategory {
		subs := &ActivateMembershipSubscriber{}
		p.register("activate_membership_subscriber", subs)
	}
	c.next.execute(p)
}

func (c *MemberAssociationChain) setNext(next Chain) {
	c.next = next
}

//end: memberAssociationChain

// upgradeAssociationChain
type UpgradeAssociationChain struct {
	next Chain
}

func (c *UpgradeAssociationChain) execute(p *payment) {
	if p.product.Category == product.MemberAssociationUpgradeCategory {
		subs := &UpgradeMembershipSubscriber{}
		p.register("upgrade_membership_subscriber", subs)
	}
	c.next.execute(p)
}

func (c *UpgradeAssociationChain) setNext(next Chain) {
	c.next = next
}

//end: upgradeAssociationChain

// sendEmailChain
type SendEmailChain struct {
	next Chain
}

func (c *SendEmailChain) execute(p *payment) {
	if p.product.Category == product.MemberAssociationCategory ||
		p.product.Category == product.MemberAssociationUpgradeCategory {
		subs := &EmailSubscriber{}
		p.register("send_email_subscriber", subs)
	}
	c.next.execute(p)
}

func (c *SendEmailChain) setNext(next Chain) {
	c.next = next
}

type PaymentComissionChain struct {
	next Chain
}

func (c *PaymentComissionChain) execute(p *payment) {
	if p.product.Category == product.PhysicalMediaCategory ||
		p.product.Category == product.BookCategory {
		subs := &PaymentComissionSubscriber{}
		p.register("payment_commission_subscriber", subs)
	}
	c.next.execute(p)
}

func (c *PaymentComissionChain) setNext(next Chain) {
	c.next = next
}

func subscribeEvents(p *payment) {
	ch1 := &SkiLearningChain{}

	ch2 := &PhysicalMediaChain{}
	ch2.setNext(ch1)

	ch3 := &BookChain{}
	ch3.setNext(ch2)

	ch4 := &MemberAssociationChain{}
	ch4.setNext(ch3)

	ch5 := &UpgradeAssociationChain{}
	ch5.setNext(ch4)

	ch6 := &SendEmailChain{}
	ch6.setNext(ch5)

	ch7 := PaymentComissionChain{}
	ch7.setNext(ch6)

	ch7.execute(p)
}
