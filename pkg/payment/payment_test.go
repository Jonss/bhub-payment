package payment

import (
	"fmt"
	"testing"

	"github.com/Jonss/bhub-payment/pkg/product"
)

func TestPaymentExecute_SkiLessonVideo(t *testing.T) {
	// setup
	payment := NewPayment(1000, product.Product{
		Name:     "Aprendendo a Esquiar",
		Category: product.VideoCategory,
	})

	if payment.packingSlip != nil {
		t.Fatal("payment.packingSlip should be nil")
	}
	// end setup

	payment.Execute()

	// assertions
	wantPackSlip := &PackingSlip{title: "adicione um vídeo gratuito de “Primeiros Socorros”"}
	wantSubscriberList := []string{"ski_learning_subscriber"}
	fmt.Println("payment", payment)
	fmt.Println("payment", payment.observers)
	fmt.Println("------------------->>>")
	if payment.packingSlip.title != wantPackSlip.title {
		t.Fatalf("packingSlip.title got %v, want %v", payment.packingSlip.title, wantPackSlip.title)
	}

	for _, s := range wantSubscriberList {
		if _, ok := payment.observers[s]; !ok {
			t.Fatalf("want payment.observer=%v, contains %v", wantSubscriberList, payment.observers)
		}
	}
}

func TestPaymentExecute_PhysicalMedia(t *testing.T) {
	// setup
	payment := NewPayment(1000, product.Product{
		Name:     "Disco dos Titãs",
		Category: product.PhysicalMediaCategory,
	})

	if payment.packingSlip != nil {
		t.Fatal("payment.packingSlip should be nil")
	}
	// end setup

	payment.Execute()

	wantPackSlip := &PackingSlip{title: "guia de remessa de envio"}
	wantSubscriberList := []string{"payment_commission_subscriber", "shipment_subscriber"}
	if payment.packingSlip.title != wantPackSlip.title {
		t.Fatalf("packingSlip.title got %v, want %v", payment.packingSlip.title, wantPackSlip.title)
	}

	fmt.Println(payment.observers)

	if len(wantSubscriberList) != len(payment.observers) {
		t.Fatalf("want len(payment.observe) %v, got %v", len(wantSubscriberList), len(payment.observers))
	}

	for _, s := range wantSubscriberList {
		if _, ok := payment.observers[s]; !ok {
			t.Fatalf("want payment.observer=%v, contains", payment.observers)
		}
	}
}

func TestPaymentExecute_activateMembership(t *testing.T) {
	// setup
	payment := NewPayment(1000, product.Product{
		Name:     "asociassao",
		Category: product.MemberAssociationCategory,
	})

	if payment.packingSlip != nil {
		t.Fatal("payment.packingSlip should be nil")
	}
	// end setup

	if payment.product.IsMember {
		t.Fatal("product membership expected false")
	}

	payment.Execute()

	// assertions
	wantSubscriberList := []string{"activate_membership_subscriber", "send_email_subscriber"}

	if !payment.product.IsMember {
		t.Fatal("product membership expected true")
	}

	if len(wantSubscriberList) != len(payment.observers) {
		t.Fatalf("want len(payment.observe) %v, got %v", len(wantSubscriberList), len(payment.observers))
	}

	for _, s := range wantSubscriberList {
		if _, ok := payment.observers[s]; !ok {
			t.Fatalf("want payment.observer=%v, contains", payment.observers)
		}
	}
}

func TestPaymentExecute_upgrateMembership(t *testing.T) {
	// setup
	payment := NewPayment(1000, product.Product{
		Name:     "upgrade_associaçao",
		Category: product.MemberAssociationUpgradeCategory,
	})

	if payment.packingSlip != nil {
		t.Fatal("payment.packingSlip should be nil")
	}
	// end setup

	if payment.product.IsMember {
		t.Fatal("product membership expected true")
	}

	payment.Execute()

	// assertions
	wantSubscriberList := []string{"upgrade_membership_subscriber", "send_email_subscriber"}

	if !payment.product.IsMember {
		t.Fatal("product membership expected true")
	}

	if len(wantSubscriberList) != len(payment.observers) {
		t.Fatalf("want len(payment.observe) %v, got %v", len(wantSubscriberList), len(payment.observers))
	}

	for _, s := range wantSubscriberList {
		if _, ok := payment.observers[s]; !ok {
			t.Fatalf("want payment.observer=%v, contains", payment.observers)
		}
	}
}

func TestPaymentExecute_bookCategory(t *testing.T) {
	// setup
	payment := NewPayment(1000, product.Product{
		Name:     "Submissao",
		Category: product.BookCategory,
	})
	// end setup

	payment.Execute()

	// assertions
	wantSubscriberList := []string{"double_packing_slip", "payment_commission_subscriber"}
	fmt.Println("obs", payment.observers)
	if len(wantSubscriberList) != len(payment.observers) {
		t.Fatalf("want len(payment.observe) %v, got %v", len(wantSubscriberList), len(payment.observers))
	}

	for _, s := range wantSubscriberList {
		if _, ok := payment.observers[s]; !ok {
			t.Fatalf("want payment.observer=%v, contains", payment.observers)
		}
	}
}
