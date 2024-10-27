package types

import "encoding/xml"

type GatewayAWebhookRequest struct {
	GatewayATransactionID string `json:"id" binding:"required"`
	UpdatedStatus         string `json:"updated_status" binding:"required"`
}

type WebhookTransactionEvent struct {
	TransactionID string `json:"transaction_id"`
	Status        string `json:"status"`
}

type SOAPEnvelope struct {
	XMLName xml.Name        `xml:"Envelope"`
	Body    WebhookSOAPBody `xml:"Body"`
}

type WebhookSOAPBody struct {
	TransactionRequest GatewayBWebhookRequest `xml:"TransactionRequest"`
}

type GatewayBWebhookRequest struct {
	TransactionID string `xml:"TransactionID"`
	Status        string `xml:"Status"`
}
