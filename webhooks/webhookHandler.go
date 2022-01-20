package webhooks

import (
	"fmt"
	"net/http"

	"github.com/stripe/stripe-go/v72"
)

func EventHandler(w http.ResponseWriter, event stripe.Event) {
	fmt.Println("\n========webhook handler============")
	switch event.Type {
	case "payment_intent.succeeded":
		paymentConfirmHandler(w, event)
	case "charge.succeeded":
		chargeHandler(w, event)
	default:
		defaultError := fmt.Sprintf("Unhandled event type: %s\n", event.Type)
		fmt.Println(defaultError)

	}
}
