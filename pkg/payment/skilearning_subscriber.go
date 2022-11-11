package payment

// ● Se o pagamento for para o vídeo específico “Aprendendo a Esquiar”, adicione um
// vídeo gratuito de “Primeiros Socorros” à guia de remessa (resultado de uma decisão
// judicial em 1997).

const learningToSki string = "Aprendendo a Esquiar"

type SkiLearningSubscriber struct {
}

func (s *SkiLearningSubscriber) update(payment *payment) {
	payment.packingSlip = &PackingSlip{title: "adicione um vídeo gratuito de “Primeiros Socorros”"}
}
