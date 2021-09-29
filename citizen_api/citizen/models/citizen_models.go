package models

import (
	"time"
)

type Address struct {
	ID            int `gorm:"primaryKey;autoIncrement"`
	CitizenId     int `gorm:"index;primaryKey"`
	PostalCode    string
	Address       string
	Number        string
	Complement    string
	Neighbourhood string
	City          string
	State         string
	Country       string
	Deleted       bool
}

type Debt struct {
	ID                  int `gorm:"primaryKey;autoIncrement"`
	DebtorID            int `gorm:"index"`
	Value               float32
	WasNegociated       bool
	CreditTakenAt       time.Time
	CreditTurnedDebitAt time.Time
	Deleted             bool
}

type Citizen struct {
	ID        int `gorm:"primaryKey;autoIncrement" json:",omitempty"`
	Name      string
	CPF       string `gorm:"index"`
	Birthdate time.Time
}
