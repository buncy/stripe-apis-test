package payment

import (
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/charge"
)

func ListCharges() []string {
	var chargeList []string

	params := &stripe.ChargeListParams{}

	i := charge.List(params)
	for i.Next() {
		c := i.Charge()
		chargeList = append(chargeList, c.ID)
	}

	return chargeList
}
