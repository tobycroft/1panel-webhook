package route

import (
	"github.com/gin-gonic/gin"
	"main.go/app"
)

func OnRoute(router *gin.Engine) {
	router.Any("/", func(context *gin.Context) {
		context.String(0, router.BasePath())
	})
	router.Any("webhook", app.WebhookController)
	router.Any("hook", app.WebhookController)
}
