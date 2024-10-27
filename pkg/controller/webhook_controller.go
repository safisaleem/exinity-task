package controller

import (
	"encoding/xml"
	constants "exinity-task/pkg/contants"
	"exinity-task/pkg/service"
	"exinity-task/pkg/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WebhookController struct {
	WebhookService service.WebhookService
}

func NewWebhookController(webhookService service.WebhookService) *WebhookController {
	return &WebhookController{WebhookService: webhookService}
}

func (wc *WebhookController) HandleGatewayAWebhook(c *gin.Context) {
	var payload types.GatewayAWebhookRequest

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var webhookTransactionEvent types.WebhookTransactionEvent
	webhookTransactionEvent.TransactionID = payload.GatewayATransactionID

	if payload.UpdatedStatus == "success" {
		webhookTransactionEvent.Status = constants.TRANSACTION_STATUS_COMPLETE
	} else {
		webhookTransactionEvent.Status = constants.TRANSACTION_STATUS_FAILED
	}

	err := wc.WebhookService.ProcessWebhook(c.Request.Context(), webhookTransactionEvent, "gatewayA")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Webhook received successfully", "data": payload})
}

func (wc *WebhookController) HandleGatewayBWebhook(c *gin.Context) {
	var soapEnvelope types.SOAPEnvelope

	// Read and parse XML body
	rawData, _ := c.GetRawData()
	if err := xml.Unmarshal(rawData, &soapEnvelope); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid XML"})
		return
	}

	// Access parsed data directly
	transactionRequest := soapEnvelope.Body.TransactionRequest

	// Now populate webhookTransactionEvent without any additional binding
	var webhookTransactionEvent types.WebhookTransactionEvent
	webhookTransactionEvent.TransactionID = transactionRequest.TransactionID

	if transactionRequest.Status == "SUCCESSFULLY_COMPLETED" {
		webhookTransactionEvent.Status = constants.TRANSACTION_STATUS_COMPLETE
	} else {
		webhookTransactionEvent.Status = constants.TRANSACTION_STATUS_FAILED
	}

	// Process the webhook event
	err := wc.WebhookService.ProcessWebhook(c.Request.Context(), webhookTransactionEvent, "gatewayB")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond in XML if required
	c.XML(http.StatusOK, gin.H{"status": "received"})
}
