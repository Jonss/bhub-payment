package payment

import (
	"testing"

	"github.com/Jonss/bhub-payment/pkg/product"
)

func TestPaymentExecute_SkiLessonVideo(t *testing.T) {

	payment := Payment{
		amount: 1000,
		product: product.Product{
			Name:     "Aprendendo a Esquiar",
			Category: product.VideoCategory,
		},
	}

	if payment.packingSlip != nil {
		t.Fatal("payment.packingSlip should be nil")
	}

	payment.Execute()

	wantPackSlip := &PackingSlip{title: "adicione um vídeo gratuito de “Primeiros Socorros”"}

	if payment.packingSlip.title != wantPackSlip.title {
		t.Fatalf("packingSlip.title got %v, want %v", payment.packingSlip.title, wantPackSlip.title)
	}
}

func TestPaymentExecute_PhysicalMedia(t *testing.T) {

	payment := Payment{
		amount: 1000,
		product: product.Product{
			Name:     "Disco dos Titãs",
			Category: product.PhysicalMediaCategory,
		},
	}

	if payment.packingSlip != nil {
		t.Fatal("payment.packingSlip should be nil")
	}

	payment.Execute()

	wantPackSlip := &PackingSlip{title: "guia de remessa de envio"}

	if payment.packingSlip.title != wantPackSlip.title {
		t.Fatalf("packingSlip.title got %v, want %v", payment.packingSlip.title, wantPackSlip.title)
	}
}
