package domain

import "time"

type Order struct {
	ID        int       // Unique identifier for the order
	Customer  string    // Name of the customer placing the order
	Items     []Item    // Slice of items in the order
	TotalCost float64   // Total cost of the order
	OrderDate time.Time // Date and time when the order was placed
}

type Item struct {
	ID    string
	Name  string
	Price float64
}

type Unit struct {
}
