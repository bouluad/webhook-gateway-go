package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gateway/internal/handler"
)

func main() {
	r := gin.Default()
	r.POST("/webhook", handler.HandleWebhook)
	r.Run(":8080")
}
