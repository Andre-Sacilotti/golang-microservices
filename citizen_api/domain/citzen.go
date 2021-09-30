package domain

import (
	"time"
)

type Address struct {
	ID            int    `gorm:"primaryKey;autoIncrement"`
	CitizenId     int    `json:"citizen_id" gorm:"index"`
	PostalCode    string `json:"postal_code" validate:"required"`
	Address       string `json:"address" validate:"required"`
	Number        string `json:"number" default:"false"`
	Complement    string `json:"complement" validate:"required"`
	Neighbourhood string `json:"neighbourhood"`
	City          string `json:"city"`
	State         string `json:"state"`
	Country       string `json:"country"`
	Deleted       bool   `json:"deleted" default:"false"`
}

type Debt struct {
	ID                  int       `gorm:"primaryKey;autoIncrement"`
	DebtorID            int       `json:"debtor_id" validate:"required" gorm:"index"`
	Value               float32   `json:"value" validate:"required"`
	WasNegociated       bool      `json:"was_negociated" default:"false"`
	CreditTakenAt       time.Time `json:"credit_taken_at" validate:"required"`
	CreditTurnedDebitAt time.Time `json:"credit_turned_debit_at"`
	Deleted             bool      `json:"deleted" default:"false"`
}

type Citizen struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:",omitempty"`
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
	InsertNewAddress(Address, int) (Address, error)
	InsertNewDebt(Debt, int) (Debt, error)
	UpdateCitizenByID(Citizen, int) (Citizen, error)
	DeleteDebt(int) (Debt, error)
	DeleteAddress(int) (Address, error)
	UpdateAddress(Address, int) (Address, error)
	UpdateDebt(Debt, int) (Debt, error)
}

type CitizenUsecase interface {
	GetCitizenByID(ID int) []Citizen
	GetCitizenByCPF(CPF string) []Citizen
	CreateCitizen(Citizen Citizen) ([]Citizen, error)
	GetDebtsByCitizenId(ID int) (res []Debt)
	GetAddressByCitizenId(ID int) (res []Address)
	GetAllCitizen() (res []Citizen)
	InsertNewAddress(Address, int) []Address
	InsertNewDebt(Debt, int) []Debt
	UpdateCitizenByID(Citizen, int) []Citizen
	DeleteDebt(int) []Debt
	DeleteAddress(int) []Address
	UpdateAddress(Address, int) []Address
	UpdateDebt(Debt, int) []Debt
}
