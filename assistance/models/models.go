package models

import (
	"time"
)

type Transaction struct {
	ID     string
	Name   string
	Type   string
	Amount float64
	Date   time.Time
}

type Product struct {
	ID       string
	Name     string
	Price    float64
	Quantity int
}
