package externalServices

import "fmt"

type NilNotifyService struct {
}

func (n *NilNotifyService) SendNotify(receiver string, message string) {

	fmt.Printf("nilNotifyService : receiver: %s , message: %s \n", receiver, message)

}

func NewNilNotifyService() *NilNotifyService {
	return &NilNotifyService{}
}
