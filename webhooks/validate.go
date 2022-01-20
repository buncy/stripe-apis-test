package webhooks

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/webhook"
)

func ValidateWebhook(webhookHandler func(http.ResponseWriter, stripe.Event)) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		fmt.Println("\n======webhook middleware============")

		stripe.Key = os.Getenv("STRIPE_SECRET_KEY")
		const MaxBodyBytes = int64(65536)
		req.Body = http.MaxBytesReader(w, req.Body, MaxBodyBytes)
		payload, err := ioutil.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading request body: %v\n", err)
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}

		// This is your Stripe CLI webhook secret.
		endpointSecret := os.Getenv("STRIPE_WEB_HOOK_KEY")

		// Pass the request body and Stripe-Signature header to ConstructEvent, along
		// with the webhook signing key.
		event, err := webhook.ConstructEvent(payload, req.Header.Get("Stripe-Signature"),
			endpointSecret)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error verifying webhook signature: %v\n", err)
			w.WriteHeader(http.StatusBadRequest) // Return a 400 error on a bad signature
			return
		}

		fmt.Println("\n======end webhook middleware============")
		//pass on the event object to the handler
		webhookHandler(w, event)
	}
}
