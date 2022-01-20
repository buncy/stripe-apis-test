package payment

import (
	"github.com/stripe/stripe-go/v72/charge"
)

func CaptureCharge(chargeId string) (string, error) {

	c, err := charge.Capture(
		chargeId,
		nil,
	)
	if err != nil {
		return "", err
	}
	return c.ID, nil
}
