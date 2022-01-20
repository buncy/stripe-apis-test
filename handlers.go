package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"stripe-test/models"
	"stripe-test/payment"

	"github.com/gorilla/mux"
)

//this is the handler for creating refund we pass a charge id as a path parameter
//and get the charge id and create a refund and return a refund id in response
func handleCreateRefund(w http.ResponseWriter, r *http.Request) {
	pars := mux.Vars(r)
	chargeId := pars["chargeId"]
	refundId, err := payment.Refund(chargeId)
	if err != nil {
		fmt.Println(err.Error())
	}
	payload := models.Response{Id: refundId}
	json.NewEncoder(w).Encode(payload)

}

//this is the handler for getting the list of charges that are made so far
// and return a the list of charge ids in response
func handleGetCharges(w http.ResponseWriter, r *http.Request) {

	charges := payment.ListCharges()

	payload := models.ListCharges{Charges: charges}
	json.NewEncoder(w).Encode(payload)

}

//this is the handler for creating charge on the default payment method of the customer
// and return a charge id in response
func handleCreateCharge(w http.ResponseWriter, r *http.Request) {

	chargeId, err := payment.CreateCharge(1000)
	if err != nil {
		fmt.Println(err.Error())
	}
	payload := models.Response{Id: chargeId}
	json.NewEncoder(w).Encode(payload)

}

//this is the handler for capturing charge previously created on a customer
// and return a charge id in response
func handleCaptureCharge(w http.ResponseWriter, r *http.Request) {
	pars := mux.Vars(r)
	chargeId := pars["chargeId"]
	chargeId, err := payment.CaptureCharge(chargeId)
	if err != nil {
		fmt.Println(err.Error())
	}
	payload := models.Response{Id: chargeId}
	json.NewEncoder(w).Encode(payload)

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "the server is running")

}
