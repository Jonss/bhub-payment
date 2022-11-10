package payment

type Observer interface {
	update(p *Payment)
}
