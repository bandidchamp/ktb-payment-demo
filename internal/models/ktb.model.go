package models

import (
	"time"

	"github.com/google/uuid"
)

type KtbPayment struct {
	Id                uuid.UUID `json:"id"`
	TransactionId     uuid.UUID `json:"transaction_id"`
	OrderId           string    `json:"order_id"`
	ChargeId          string    `json:"charge_id"`
	SourceType        string    `json:"source_type"`
	QrId              string    `json:"qr_id"`
	Mode              string    `json:"mode"`
	Token             string    `json:"token"`
	Mid               string    `json:"mid"`
	Tid               string    `json:"tid"`
	SmartpayId        string    `json:"smartpay_id"`
	Term              int       `json:"term"`
	Qty               int       `json:"qty"`
	Amount            float64   `json:"amount"`
	Description       string    `json:"description"`
	Currency          string    `json:"currency"`
	DccCurrency       string    `json:"dcc_currency"`
	ReferenceOrder    string    `json:"reference_order"`
	CustomerId        string    `json:"customer_id"`
	Transaction_state string    `json:"transaction_state"`
	RedirectUrl       string    `json:"redirect_url"`
	Status            string    `json:"status"`
	FailureCode       string    `json:"failure_code"`
	FailureMessage    string    `json:"failure_message"`
	CreatedAt         time.Time `json:"created_at"`
	PaidAt            time.Time `json:"paid_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type KtbCrateQROrderRequest struct {
	Amount         float64    `json:"amount"`
	Currency       string     `json:"currency"`
	Description    string     `json:"description"`
	SourceType     string     `json:"source_type"`
	ReferenceOrder string     `json:"reference_order"`
	Ref1           string     `json:"ref_1"`
	Metadata       []metadata `json:"metadata"`
	Customer       []customer `json:"customer"`
}

type KtbCrateQROrderResponse struct {
	Id                  string         `json:"id"`
	Object              string         `json:"object"`
	Created             string         `json:"created"`
	Livemode            string         `json:"livemode"`
	Amount              string         `json:"amount"`
	Currency            string         `json:"currency"`
	Customer            []customer     `json:"customer"`
	Description         string         `json:"description"`
	Metadata            []metadata     `json:"metadata"`
	Status              string         `json:"status"`
	Reference_order     string         `json:"reference_order"`
	Source_type         string         `json:"source_type"`
	Additional_data     additionalData `json:"additional_data"`
	Failure_code        string         `json:"failure_code"`
	Failure_message     string         `json:"failure_message"`
	Expire_time_seconds string         `json:"expire_time_seconds"`
}

type metadata struct {
	Item   string  `json:"item"`
	Qty    int     `json:"qty"`
	Amount float64 `json:"amount"`
}

type customer struct {
	CustomerId string `json:"customer_id"`
}

type additionalData struct {
	AdditionalData string `json:"additional_data"`
}
