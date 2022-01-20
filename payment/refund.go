package payment

import (
	"fmt"

	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/refund"
)

func Refund(chargeId string) (string, error) {
	refundParams := &stripe.RefundParams{
		Charge: &chargeId,
	}
	r, err := refund.New(refundParams)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	fmt.Println(r.ID)
	return r.ID, nil
}
