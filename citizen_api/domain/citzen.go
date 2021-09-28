package domain

import (
	"time"
)

type Address struct {
	CitizenId     int    `json:"citizen_id" gorm:"index;primaryKey"`
	PostalCode    string `json:"postal_code" validate:"required"`
	Address       string `json:"address" validate:"required"`
	Number        string `json:"number" default:"false"`
	Complement    string `json:"complement" validate:"required"`
	Neighbourhood string `json:"neighbourhood"`
	City          string `json:"city"`
	State         string `json:"state"`
	Country       string `json:"country"`
}

type Debt struct {
	DebtorID            int       `json:"debtor_id" validate:"required" gorm:"index;primaryKey"`
	Value               float32   `json:"value" validate:"required"`
	WasNegociated       bool      `json:"was_negociated" default:"false"`
	CreditTakenAt       time.Time `json:"credit_taken_at" validate:"required"`
	CreditTurnedDebitAt time.Time `json:"credit_turned_debit_at"`
}

type Citizen struct {
	ID        int       `gorm:"primaryKey"`
	Name      string    `json:"password" validate:"required" gorm:"index"`
	CPF       string    `json:"cpf" validate:"required"`
	Birthdate time.Time `json:"birthdate" validate:"required"`
	Debts     []Debt    `json:"debts" validate:"required"`
	Address   []Address `json:"address" validate:"required"`
}

type CitizenRepository interface {
	GetCitizenByID(ID int) (Citizen, error)
	GetCitizenByCPF(CPF string) (Citizen, error)
	CreateCitizen(Citizen Citizen) (Citizen, error)
	GetDebtsByCitizenId(ID int) (res []Debt, err error)
	GetAddressByCitizenId(ID int) (res []Address, err error)
	GetAllCitizen() (res []Citizen, err error)
}

type CitizenUsecase interface {
	GetCitizenByID(ID int) (Citizen, error)
	GetCitizenByCPF(CPF string) (Citizen, error)
}
