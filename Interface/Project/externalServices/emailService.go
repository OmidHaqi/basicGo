package externalServices

import (
	"fmt"
	"notification/entities"
)

type EmailService struct {
}

func (e *EmailService) SendMessage(order *entities.Order) {
	fmt.Printf("Email send: %v\n", order)
}

func (e *EmailService) SendNotify(receiver string, message string) {
	fmt.Printf("Email send : receiver : %s , message : %s\n", receiver, message)
}

func NewEmailService() *EmailService {
	return &EmailService{}
}
