package repository

import (
	"errors"

	"github.com/Andre-Sacilotti/golang-credit-backend/citizen_api/citizen/models"
	"github.com/Andre-Sacilotti/golang-credit-backend/citizen_api/domain"
	"github.com/Andre-Sacilotti/golang-credit-backend/citizen_api/utils"
	"gorm.io/gorm"
)

type mysqlCitizenRepository struct {
	Conn *gorm.DB
}

func CitizenRepositoryInterface(Conn *gorm.DB) domain.CitizenRepository {
	return &mysqlCitizenRepository{Conn}
}

func (CitizenRepo *mysqlCitizenRepository) GetDebtsByCitizenId(ID int) (res []domain.Debt, err error) {
	var debts []domain.Debt
	var dcrypted_debts []domain.Debt

	if result := CitizenRepo.Conn.First(&debts, "debtor_id = ?", ID); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return debts, domain.ErrNotFound
		}
		return
	}

	for _, element := range debts {
		tmp_debt := domain.Debt{ID: element.ID, DebtorID: element.DebtorID, Value: element.Value,
			WasNegociated: element.WasNegociated, CreditTakenAt: element.CreditTakenAt,
			CreditTurnedDebitAt: element.CreditTurnedDebitAt,
		}
		dcrypted_debts = append(dcrypted_debts, tmp_debt)

	}

	return dcrypted_debts, err

}

func (CitizenRepo *mysqlCitizenRepository) GetAddressByCitizenId(ID int) (res []domain.Address, err error) {
	var address []domain.Address
	var decrypted_address []domain.Address

	if result := CitizenRepo.Conn.First(&address, "citizen_id = ?", ID); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return address, domain.ErrNotFound
		}
		return
	}

	for _, element := range address {
		tmp_address := domain.Address{ID: element.ID, CitizenId: element.CitizenId, PostalCode: utils.Decrypt(element.PostalCode),
			Address: utils.Decrypt(element.Address), Number: utils.Decrypt(element.Number),
			Complement: utils.Decrypt(element.Complement), Neighbourhood: utils.Decrypt(element.Neighbourhood),
			City: utils.Decrypt(element.City), State: utils.Decrypt(element.State), Country: utils.Decrypt(element.Country),
		}

		decrypted_address = append(decrypted_address, tmp_address)
	}

	return decrypted_address, err

}

func (CitizenRepo *mysqlCitizenRepository) GetCitizenByID(ID int) (res domain.Citizen, err error) {
	var citizen models.Citizen

	if result := CitizenRepo.Conn.First(&citizen, "id = ?", ID); result.Error != nil {

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return_citizen := domain.Citizen{ID: citizen.ID, Name: citizen.Name, CPF: citizen.CPF, Birthdate: citizen.Birthdate}
			return return_citizen, domain.ErrNotFound
		}
		return
	}

	debts, _ := CitizenRepo.GetDebtsByCitizenId(ID)
	address, _ := CitizenRepo.GetAddressByCitizenId(ID)

	return_citizen := domain.Citizen{
		Name: utils.Decrypt(citizen.Name), CPF: utils.Decrypt(citizen.CPF), ID: citizen.ID,
		Birthdate: citizen.Birthdate, Debts: debts, Address: address,
	}
	return return_citizen, err

}

func (CitizenRepo *mysqlCitizenRepository) GetCitizenByCPF(CPF string) (res domain.Citizen, err error) {
	var citizen models.Citizen

	if result := CitizenRepo.Conn.First(&citizen, "cpf = ?", utils.Encrypt(CPF)); result.Error != nil {

		return_citizen := domain.Citizen{
			Name: citizen.Name, CPF: citizen.CPF,
			Birthdate: citizen.Birthdate,
		}

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {

			return return_citizen, domain.ErrNotFound
		}
		return return_citizen, domain.ErrNotFound
	}
	debts, _ := CitizenRepo.GetDebtsByCitizenId(citizen.ID)
	address, _ := CitizenRepo.GetAddressByCitizenId(citizen.ID)

	return_citizen := domain.Citizen{
		Name: citizen.Name, CPF: citizen.CPF, ID: citizen.ID,
		Birthdate: citizen.Birthdate, Debts: debts, Address: address,
	}
	return return_citizen, err
}

func (CitizenRepo *mysqlCitizenRepository) GetAllCitizen() (res []domain.Citizen, err error) {
	citizens_model := []models.Citizen{}
	citizens := []domain.Citizen{}

	result := CitizenRepo.Conn.Find(&citizens_model)

	for _, element := range citizens_model {

		tmp_citizen := domain.Citizen{ID: element.ID, Name: utils.Decrypt(element.Name), CPF: utils.Decrypt(element.CPF), Birthdate: element.Birthdate}
		citizens = append(citizens, tmp_citizen)
	}

	return citizens, result.Error
}

func (CitizenRepo *mysqlCitizenRepository) InsertNewAddress(Address domain.Address, CitizenId int) (res domain.Address, err error) {
	AddressModel := models.Address{CitizenId: CitizenId, PostalCode: utils.Encrypt(Address.PostalCode),
		Address: utils.Encrypt(Address.Address), Number: utils.Encrypt(Address.Number),
		Complement: utils.Encrypt(Address.Complement), Neighbourhood: utils.Encrypt(Address.Neighbourhood),
		City: utils.Encrypt(Address.City), State: utils.Encrypt(Address.State), Country: utils.Encrypt(Address.Country),
	}

	if result := CitizenRepo.Conn.Create(&AddressModel); result.Error != nil {

		return Address, result.Error
	}

	return Address, err

}

func (CitizenRepo *mysqlCitizenRepository) InsertNewDebt(Debt domain.Debt, CitizenId int) (res domain.Debt, err error) {
	DebtModel := models.Debt{DebtorID: CitizenId, Value: Debt.Value,
		WasNegociated: Debt.WasNegociated, CreditTakenAt: Debt.CreditTakenAt,
		CreditTurnedDebitAt: Debt.CreditTurnedDebitAt,
	}

	if result := CitizenRepo.Conn.Create(&DebtModel); result.Error != nil {

		return Debt, result.Error
	}

	return Debt, err

}

func (CitizenRepo *mysqlCitizenRepository) CreateCitizen(citizen domain.Citizen) (res domain.Citizen, err error) {
	citizen_model := models.Citizen{Name: utils.Encrypt(citizen.Name),
		CPF: utils.Encrypt(citizen.CPF), Birthdate: citizen.Birthdate}

	if result := CitizenRepo.Conn.Create(&citizen_model); result.Error != nil {

		return citizen, result.Error
	}

	for _, element := range citizen.Address {
		CitizenRepo.InsertNewAddress(element, citizen_model.ID)
	}

	for _, element := range citizen.Debts {
		CitizenRepo.InsertNewDebt(element, citizen_model.ID)
	}

	return_citizen := domain.Citizen{ID: citizen_model.ID, Name: citizen.Name, CPF: citizen.CPF, Birthdate: citizen.Birthdate}
	return return_citizen, err
}

func (CitizenRepo *mysqlCitizenRepository) UpdateCitizenByID(Citizen domain.Citizen) (res domain.Citizen, err error) {
	CitizenWithId, _ := CitizenRepo.GetCitizenByCPF(Citizen.CPF)
	CitizenOldDebts, _ := CitizenRepo.GetDebtsByCitizenId(CitizenWithId.ID)
	CitizenOldAddress, _ := CitizenRepo.GetAddressByCitizenId(CitizenWithId.ID)

	for _, OldDebt := range CitizenOldDebts {
		DebtExists := false
		for _, Debt := range Citizen.Debts {
			if (OldDebt.Value == Debt.Value) && (OldDebt.CreditTakenAt == Debt.CreditTakenAt) && (OldDebt.ID == Debt.ID) {
				DebtExists = true
			}
		}

		if !DebtExists {
			CitizenRepo.Conn.Model(&OldDebt).Update("deleted", true)
		}

	}

	for _, Debt := range Citizen.Debts {
		DebtExists := false
		for _, OldDebt := range CitizenOldDebts {
			if (OldDebt.Value == Debt.Value) && (OldDebt.CreditTakenAt == Debt.CreditTakenAt) && (OldDebt.ID == Debt.ID) {
				DebtExists = true
			}
		}

		if !DebtExists {
			CitizenRepo.InsertNewDebt(Debt, CitizenWithId.ID)
		}

	}

	for _, OldAddress := range CitizenOldAddress {
		AddressExists := false
		for _, Address := range Citizen.Address {
			if (OldAddress.PostalCode == Address.PostalCode) && (OldAddress.Number == Address.Number) && (OldAddress.ID == Address.ID) {
				AddressExists = true
			}
		}

		if !AddressExists {
			CitizenRepo.Conn.Model(&OldAddress).Update("deleted", true)
		}
	}

	for _, Address := range Citizen.Address {
		AddressExists := false
		for _, OldAddress := range CitizenOldAddress {
			if (OldAddress.PostalCode == Address.PostalCode) && (OldAddress.Number == Address.Number) && (OldAddress.ID == Address.ID) {
				AddressExists = true
			}
		}

		if !AddressExists {
			CitizenRepo.InsertNewAddress(Address, CitizenWithId.ID)
		}

	}

	CitizenRepo.Conn.Model(&CitizenWithId).Update("name", utils.Encrypt(Citizen.Name))
	CitizenRepo.Conn.Model(&CitizenWithId).Update("cpf", utils.Encrypt(Citizen.CPF))
	CitizenRepo.Conn.Model(&CitizenWithId).Update("birthdate", Citizen.Birthdate)

	return Citizen, err

}

func (CitizenRepo *mysqlCitizenRepository) DeleteDebt(ID int) (res domain.Debt, err error) {
	var Debt models.Debt

	if result := CitizenRepo.Conn.First(&Debt, "id = ?", ID); result.Error != nil {

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return res, domain.ErrNotFound
		}
		return
	}
	CitizenRepo.Conn.Model(&Debt).Update("deleted", true)
	TmpDebt := domain.Debt{ID: Debt.ID, DebtorID: Debt.DebtorID, Value: Debt.Value,
		WasNegociated: Debt.WasNegociated, CreditTakenAt: Debt.CreditTakenAt,
		CreditTurnedDebitAt: Debt.CreditTurnedDebitAt, Deleted: true,
	}
	return TmpDebt, err

}

// func (CitizenRepo *mysqlCitizenRepository) UpdateDebt(ID int) (res domain.Debt, err error) {
// 	return
// }

func (CitizenRepo *mysqlCitizenRepository) DeleteAddress(ID int) (res domain.Address, err error) {
	var Address models.Address

	if result := CitizenRepo.Conn.First(&Address, "id = ?", ID); result.Error != nil {

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return res, domain.ErrNotFound
		}
		return
	}
	CitizenRepo.Conn.Model(&Address).Update("deleted", true)
	TmpDebt := domain.Address{CitizenId: Address.ID, PostalCode: utils.Decrypt(Address.PostalCode),
		Address: utils.Decrypt(Address.Address), Number: utils.Decrypt(Address.Number),
		Complement: utils.Decrypt(Address.Complement), Neighbourhood: utils.Decrypt(Address.Neighbourhood),
		City: utils.Decrypt(Address.City), State: utils.Decrypt(Address.State), Country: utils.Decrypt(Address.Country),
		Deleted: true,
	}
	return TmpDebt, err
}

// func (CitizenRepo *mysqlCitizenRepository) UpdateAddress(ID int) (res domain.Citizen, err error) {
// 	return
// }
