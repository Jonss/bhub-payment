package payment

import (
	"testing"

	"github.com/Jonss/bhub-payment/pkg/product"
)

func TestPaymentExecute_SkiLessonVideo(t *testing.T) {

	testCases := []struct {
		name         string
		payment      Payment
		wantProducts []product.Product
	}{
		{
			name: "should add first aid video when product is ski lessons",
			payment: Payment{
				amount: 1000,
				products: []product.Product{
					{
						Name:     "Aprendendo a Esquiar",
						Category: product.VideoCategory,
					},
				},
			},
			wantProducts: []product.Product{
				{
					Name:     "Aprendendo a Esquiar",
					Category: product.VideoCategory,
				},
				{
					Category: product.VideoCategory,
					Name:     "Primeiros Socorros",
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.payment.Execute()

			for i := 0; i < len(tc.wantProducts); i++ {
				want := tc.wantProducts[i]
				got := tc.payment.products[i]
				if got.Name != want.Name {
					t.Errorf("Payment.product.name, got %v, want %v", got.Name, want.Name)
				}

				if string(want.Category) != string(got.Category) {
					t.Errorf("Payment.product.category, got %v, want %v", got.Category, want.Category)
				}
			}

		})
	}

}
