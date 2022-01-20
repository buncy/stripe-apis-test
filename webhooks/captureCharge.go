package webhooks

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/stripe/stripe-go/v72"
)

func chargeHandler(w http.ResponseWriter, event stripe.Event) {
	fmt.Println("\n========= entered payment confirm webhook==============")

	var charge stripe.Charge
	err := json.Unmarshal(event.Data.Raw, &charge)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing webhook JSON: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Printf("\n chargeId : %s successful\n", charge.ID)

	w.WriteHeader(http.StatusOK)
}
