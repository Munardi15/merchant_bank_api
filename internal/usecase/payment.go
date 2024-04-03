package usecase

import (
	"errors"
	"fmt"
)

type PaymentUseCase interface {
	ProcessPayment(sessionID string, amount int) error
}

type paymentUseCase struct {
}

func NewPaymentUseCase() PaymentUseCase {
	return &paymentUseCase{}
}

func (uc *paymentUseCase) ProcessPayment(sessionID string, amount int) error {
	if sessionID == "" {
		return errors.New("session ID is required")
	}

	if amount <= 0 {
		return errors.New("invalid payment amount")
	}

	fmt.Printf("Processing payment for session ID %s, amount: %d\n", sessionID, amount)

	return nil
}
