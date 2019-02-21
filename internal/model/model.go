package model

import (
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Code        string
	Description string
	Image       string
	Price       int
}

type PaymentOrder struct {
	gorm.Model
	CardMask       string
	No             string
	Description    string
	Amount         int
	RefundedAmount int
	Status         string
}

type SaveCard struct {
	gorm.Model
	CardMask string
	Token    string
}

type Card struct {
	Number   string
	ExpMonth string
	ExpYear  string
	Cvv      string
}
