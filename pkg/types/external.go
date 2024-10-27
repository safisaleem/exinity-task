package types

import "encoding/xml"

// gatewayA structs
type GatewayATransactionRequest struct {
	Amount float64 `json:"amount"`
	UserID string  `json:"user_id"`
	Type   string  `json:"type"`
}

type GatewayATransactionResponse struct {
	TransactionID string `json:"transaction_id"`
	Status        string `json:"status"`
}

// gatewayB structs
type GatewayBTransactionRequest struct {
	XMLName xml.Name `xml:"soapenv:Envelope"`
	Xmlns   string   `xml:"xmlns:soapenv,attr"`
	Web     string   `xml:"xmlns:web,attr"`
	Body    SOAPBody `xml:"soapenv:Body"`
}

type SOAPBody struct {
	Transaction GatewayBRawTransactionRequest `xml:"web:TransactionRequest"`
}

type GatewayBRawTransactionRequest struct {
	Amount float64 `xml:"web:Amount"`
	UserID string  `xml:"web:UserID"`
	Type   string  `xml:"web:Type"`
}

type GatewayBTransactionResponse struct {
	XMLName xml.Name                        `xml:"Envelope"`
	Body    GatewayBTransactionResponseBody `xml:"Body"`
}

type GatewayBTransactionResponseBody struct {
	TransactionResponse GatewayBRawTransactionResponse `xml:"TransactionResponse"`
}

type GatewayBRawTransactionResponse struct {
	TransactionID string `xml:"TransactionID"`
	Status        string `xml:"Status"`
}
