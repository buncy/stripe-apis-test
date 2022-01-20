package payment

import (
	"fmt"
	"net/http"

	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/charge"
	"github.com/stripe/stripe-go/v72/paymentintent"
	"github.com/stripe/stripe-go/v72/paymentmethod"
)

func CreateCheckoutSession(w http.ResponseWriter, r *http.Request) {

	cardparams := &stripe.PaymentMethodListParams{
		Customer: stripe.String("cus_KyoozgMeI7xmUx"),
		Type:     stripe.String("card"),
	}
	result := paymentmethod.List(cardparams)
	fmt.Println(result)
	params := &stripe.PaymentIntentParams{
		Amount:        stripe.Int64(1099),
		Currency:      stripe.String(string(stripe.CurrencyUSD)),
		Customer:      stripe.String("cus_KyoozgMeI7xmUx"),
		PaymentMethod: stripe.String("card_1KIrBB2Ri1hXKaRW1lzx1zPe"),
		Confirm:       stripe.Bool(true),
		OffSession:    stripe.Bool(true),
	}

	_, err := paymentintent.New(params)

	if err != nil {
		if stripeErr, ok := err.(*stripe.Error); ok {
			// Error code will be authentication_required if authentication is needed
			fmt.Printf("Error code: %v", stripeErr.Code)

			paymentIntentID := stripeErr.PaymentIntent.ID
			paymentIntent, _ := paymentintent.Get(paymentIntentID, nil)

			fmt.Printf("PI: %v", paymentIntent.ID)
		}
	}
}

func CreateCharge(amount int64) (string, error) {

	params := &stripe.ChargeParams{
		Amount:      stripe.Int64(amount),
		Currency:    stripe.String(string(stripe.CurrencyUSD)),
		Description: stripe.String("Test charge for Chai_test"),
		//   Source: &stripe.SourceParams{Token: stripe.String("tok_amex")},
		Customer: stripe.String("cus_KyoozgMeI7xmUx"),
		Capture:  stripe.Bool(false),
	}
	c, err := charge.New(params)
	if err != nil {
		return "", err
	}
	fmt.Printf("\nthis is the charge ID : %s and amount : %d \n", c.ID, amount)
	return c.ID, nil
}
