package models

import (
	"time"
)

type Address struct {
	CitizenId     int `gorm:"index;primaryKey"`
	PostalCode    string
	Address       string
	Number        string
	Complement    string
	Neighbourhood string
	City          string
	State         string
	Country       string
}

type Debt struct {
	DebtorID            int `gorm:"index;primaryKey"`
	Value               float32
	WasNegociated       bool
	CreditTakenAt       time.Time
	CreditTurnedDebitAt time.Time
}

type Citizen struct {
	ID        int `gorm:"primaryKey;autoIncrement" json:",omitempty"`
	Name      string
	CPF       string `gorm:"index"`
	Birthdate time.Time
}
