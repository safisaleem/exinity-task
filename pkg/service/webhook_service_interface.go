package service

import (
	"context"
	"exinity-task/pkg/types"
)

type WebhookService interface {
	ProcessWebhook(context context.Context, webhookTransactionEvent types.WebhookTransactionEvent, gatewayType string) error
}
