package http

import (
	"net/http"

	"merchant_bank_api/internal/usecase"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	authUseCase    usecase.AuthUseCase
	paymentUseCase usecase.PaymentUseCase
}

func NewHandler(authUC usecase.AuthUseCase, paymentUC usecase.PaymentUseCase) *Handler {
	return &Handler{
		authUseCase:    authUC,
		paymentUseCase: paymentUC,
	}
}

func (h *Handler) Login(c *gin.Context) {
	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sessionID, err := h.authUseCase.Login(loginData.Username, loginData.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"session_id": sessionID})
}

func (h *Handler) Payment(c *gin.Context) {
	var paymentData struct {
		SessionID string `json:"session_id"`
		Amount    int    `json:"amount"`
	}
	if err := c.ShouldBindJSON(&paymentData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.paymentUseCase.ProcessPayment(paymentData.SessionID, paymentData.Amount); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Payment successful"})
}

func (h *Handler) Logout(c *gin.Context) {
	var logoutData struct {
		SessionID string `json:"session_id"`
	}
	if err := c.ShouldBindJSON(&logoutData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.authUseCase.Logout(logoutData.SessionID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}
