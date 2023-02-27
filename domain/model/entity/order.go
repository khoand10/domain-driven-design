package entity

import "time"

type Order struct {
	ID              string
	CustomerID      string
	ProductID       string
	Amount          int
	CreatedDate     time.Time
	DeliveryDate    time.Time
	DeliveryAddress string
}
