package entities

type Order struct {
	ID               int64
	Name             string
	UserID           string
	Price            float64
	Status           bool
	NotificationType NotificationType
}
