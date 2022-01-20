package models

type Response struct {
	Id string `json:"responseId"`
}

type ListCharges struct {
	Charges []string `json:"charges"`
}
