package main

import (
	"notification/entities"
	"notification/services"
)

func main() {

	order := entities.Order{
		ID:               1,
		Name:             "Umut",
		UserID:           "989100090708",
		Price:            float64(100),
		Status:           true,
		NotificationType: entities.SMS,
	}

	orderService := services.NewOrderService()
	orderService.CreateOrder(&order)

	order1 := entities.Order{
		ID:               2,
		Name:             "Omid",
		UserID:           "989100090708",
		Price:            float64(100),
		Status:           true,
		NotificationType: entities.SMS,
	}
	order2 := entities.Order{
		ID:               2,
		Name:             "Omid",
		UserID:           "989100090708",
		Price:            float64(100),
		Status:           true,
		NotificationType: "",
	}

	orderService = services.NewOrderService()
	orderService.CreateOrder(&order1)
	orderService.CreateOrder(&order2)

}
