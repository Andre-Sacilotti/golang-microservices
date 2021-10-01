package domain

import (
	"time"
)

type Address struct {
	ID            int    `gorm:"primaryKey;autoIncrement" json:",omitempty"`
	CitizenId     int    `json:"citizen_id,omitempty" gorm:"index"`
	PostalCode    string `json:"postal_code,omitempty" `
	Address       string `json:"address,omitempty" `
	Number        string `json:"number,omitempty" default:"false"`
	Complement    string `json:"complement,omitempty" `
	Neighbourhood string `json:"neighbourhood,omitempty"`
	City          string `json:"city,omitempty"`
	State         string `json:"state,omitempty"`
	Country       string `json:"country,omitempty"`
	Deleted       bool   `json:"deleted,omitempty" default:"false"`
}

type Debt struct {
	ID                  int       `gorm:"primaryKey;autoIncrement" json:",omitempty"`
	DebtorID            int       `json:"debtor_id,omitempty"  gorm:"index"`
	Value               float32   `json:"value,omitempty" `
	WasNegociated       bool      `json:"was_negociated,omitempty" default:"false"`
	CreditTakenAt       time.Time `json:"credit_taken_at,omitempty" `
	CreditTurnedDebitAt time.Time `json:"credit_turned_debit_at,omitempty"`
	Deleted             bool      `json:"deleted,omitempty" default:"false"`
}

type Citizen struct {
	ID      int       `gorm:"primaryKey;autoIncrement" json:",omitempty"`
	Name    string    `json:"name,omitempty"  gorm:"index"`
	CPF     string    `json:"cpf,omitempty"`
	Debts   []Debt    `json:"debts,omitempty"`
	Address []Address `json:"address,omitempty" `
}

type CitizenRepository interface {
	GetCitizenByCPF(CPF string) (Citizen, error)
	CreateCitizen(Citizen Citizen) (Citizen, error)
	GetDebtsByCitizenCPF(CPF string) (res []Debt, err error)
	GetAddressByCitizenCPF(CPF string) (res []Address, err error)
	GetAllCitizen(offset int, limit int) (res []Citizen, err error)
	InsertNewAddress(Address, string) (Address, error)
	InsertNewDebt(Debt, string) (Debt, error)
	UpdateCitizenByCPF(Citizen, string) (Citizen, error)
	DeleteDebt(int) (Debt, error)
	DeleteAddress(int) (Address, error)
	UpdateAddress(Address, int) (Address, error)
	UpdateDebt(Debt, int) (Debt, error)
}

type CitizenUsecase interface {
	GetCitizenByCPF(CPF string) []Citizen
	CreateCitizen(Citizen Citizen) ([]Citizen, error)
	GetAllCitizen(offset int, limit int) (res []Citizen)
	UpdateCitizenByCPF(Citizen, string) []Citizen

	GetDebtsByCitizenCPF(CPF string) (res []Debt)
	InsertNewDebt(Debt, string) ([]Debt, error)
	UpdateDebt(Debt, int) ([]Debt, error)
	DeleteDebt(int) []Debt

	GetAddressByCitizenCPF(CPF string) (res []Address)
	InsertNewAddress(Address, string) ([]Address, error)
	DeleteAddress(int) []Address
	UpdateAddress(Address, int) ([]Address, error)
}
