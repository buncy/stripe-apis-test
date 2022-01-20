package webhooks

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/stripe/stripe-go/v72"
)

func paymentConfirmHandler(w http.ResponseWriter, event stripe.Event) {
	fmt.Println("\n========= entered payment confirm webhook==============")

	var paymentIntent stripe.PaymentIntent
	err := json.Unmarshal(event.Data.Raw, &paymentIntent)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing webhook JSON: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//we can update a database record here coresponding to the paymentIntent_id and userId
	//for now just logging the confirmation

	fmt.Printf("\n payment recieved for payment intent %s and customer id %s \n", paymentIntent.ID, paymentIntent.Customer.ID)

	fmt.Println("PaymentIntent was successful!")

	w.WriteHeader(http.StatusOK)
}
