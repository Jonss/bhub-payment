package payment

import (
	"fmt"

	"github.com/Jonss/bhub-payment/pkg/product"
)

// ● Se o pagamento for para o vídeo específico “Aprendendo a Esquiar”, adicione um
// vídeo gratuito de “Primeiros Socorros” à guia de remessa (resultado de uma decisão
// judicial em 1997).

const learningToSki string = "Aprendendo a Esquiar"

type SkiLearningSubscriber struct {
}

func (s *SkiLearningSubscriber) update(payment *Payment) {
	fmt.Println("adicionando video gratuito de primeiros socorros")
	payment.products = append(payment.products, firstAidVideo)
}

var firstAidVideo = product.Product{
	Category: product.VideoCategory,
	Name:     "Primeiros Socorros",
}
