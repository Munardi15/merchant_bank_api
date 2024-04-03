package main

import (
	"merchant_bank_api/internal/delivery/http"
	"merchant_bank_api/internal/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Initialize use cases
	authUseCase := usecase.NewAuthUseCase()
	paymentUseCase := usecase.NewPaymentUseCase()

	// Initialize HTTP handlers
	handler := http.NewHandler(authUseCase, paymentUseCase)

	// Register HTTP routes
	api := r.Group("/api")
	{
		api.POST("/login", handler.Login)
		api.POST("/payment", handler.Payment)
		api.POST("/logout", handler.Logout)
	}

	// Register middleware
	api.Use(http.AuthMiddleware())

	r.Run(":8080")
}
