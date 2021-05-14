package models

import "github.com/PayLaterCLI/constants"

type DataTransportRequest struct {
	Method       constants.Method `json:"api"`
	UserData     *User            `json:"userData,omitempty"`
	MerchantData *Merchant        `json:"merchantData,omitempty"`
	Transaction  *Transaction     `json:"transaction,omitempty"`
	Reports      *Report          `json:"reports"`
}

type DataTransportResponse struct {
	Method  constants.Method `json:"api"`
	Code    constants.Code   `json:"code"`
	Error   string           `json:"error,omitempty"`
	Message string           `json:"Message,omitempty"`
	Data    []string         `json:"data,omitempty"`
}

type Transaction struct {
	User     string  `json:"user"`
	Merchant string  `json:"merchant"`
	Amount   float64 `json:"amount"`
}

type Report struct {
	User               string `json:"user"`
	Merchant           string `json:"merchant"`
	UsersAtCreditLimit bool   `json:"users_at_credit_limit"`
	TotalDues          bool   `json:"total_dues"`
}
