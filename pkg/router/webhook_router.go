package router

import (
	"exinity-task/pkg/controller"

	"github.com/gin-gonic/gin"
)

func InitWebhookRoutes(router *gin.Engine, webhookController *controller.WebhookController) {
	webhookRouter := router.Group("/webhook")

	webhookRouter.POST("/gatewayA", webhookController.HandleGatewayAWebhook)
	webhookRouter.POST("/gatewayB", webhookController.HandleGatewayBWebhook)
}
