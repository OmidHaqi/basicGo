package externalServices

import "notification/entities"

type Notifier interface {
	SendNotify(receiver string, message string)
}

func NewNotifier(notificationType entities.NotificationType) Notifier {

	switch notificationType {
	case entities.SMS:
		return NewSmsService()

	case entities.EMAIL:
		return NewEmailService()

	default:
		return NewNilNotifyService()
	}

}
