package domain

import (
	"time"
)

type Address struct {
	Id            int    `json:"id" validate:"required"`
	PostalCode    string `json:"postal_code" validate:"required"`
	Address       string `json:"address" validate:"required"`
	Number        string `json:"number" default:"false"`
	Complement    string `json:"complement" validate:"required"`
	Neighbourhood string `json:"neighbourhood"`
	City          string `json:"city"`
	State         string `json:"state"`
	Country       string `json:"country"`
	CitizenId     int    `json:"citizen_id"`
}

type Debt struct {
	Id                  int       `json:"id" validate:"required"`
	DebtorID            int       `json:"debtor_id" validate:"required"`
	Value               float32   `json:"value" validate:"required"`
	WasNegociated       bool      `json:"was_negociated" default:"false"`
	CreditTakenAt       time.Time `json:"credit_taken_at" validate:"required"`
	CreditTurnedDebitAt time.Time `json:"credit_turned_debit_at"`
}

type Citizen struct {
	Id        string    `json:"user" validate:"required"`
	Name      string    `json:"password" validate:"required"`
	CPF       string    `json:"cpf" validate:"required"`
	birthdate string    `json:"birthdate" validate:"required"`
	debts     []Debt    `json:"debts"`
	address   []Address `json:"adress"`
}
