package externalServices

import (
	"fmt"
	"notification/entities"
)

type SmsService struct {
}

func (s *SmsService) SendMessage(order *entities.Order) {
	fmt.Printf("Sms send: %v\n", order)
}

func (s *SmsService) SendNotify(receiver string, message string) {
	fmt.Printf("sms send : receiver : %s , message : %s\n", receiver, message)
}

func NewSmsService() *SmsService {
	return &SmsService{}
}
