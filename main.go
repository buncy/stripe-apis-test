package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/stripe/stripe-go/v72"
)

func main() {
	//ideally we would never have our API keys hard coded inside our code
	//we would either get them as an Environment variable or from a .env file
	//since this is only a test key
	stripe.Key = "sk_test_51K67xD2Ri1hXKaRW88YXXsi1x9YyOp0TaKSoX1h5TYcG9fR1WuWAdijipgKfNcujJtqxrmsY6pI3Obb5LnUXl2Mh00HMvsgVte"
	r := mux.NewRouter()
	r.Use(loggingMiddleware)
	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/api/v1/create_charge", handleCreateCharge).Methods("GET")
	r.HandleFunc("/api/v1/capture_charge/{chargeId}", handleCaptureCharge).Methods("GET")
	r.HandleFunc("/api/v1/create_refund/{chargeId}", handleCreateRefund).Methods("GET")
	r.HandleFunc("/api/v1/get_charges", handleGetCharges).Methods("GET")
	http.ListenAndServe(":9040", r)
}
