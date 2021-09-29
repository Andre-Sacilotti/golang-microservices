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

	if result := CitizenRepo.Conn.First(&debts, "debtor_id = ?", ID); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return debts, domain.ErrNotFound
		}
		return
	}
	return debts, err

}

func (CitizenRepo *mysqlCitizenRepository) GetAddressByCitizenId(ID int) (res []domain.Address, err error) {
	var debts []domain.Address

	if result := CitizenRepo.Conn.First(&debts, "citizen_id = ?", ID); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return debts, domain.ErrNotFound
		}
		return
	}
	return debts, err

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

	if result := CitizenRepo.Conn.First(&citizen, "cpf = ?", CPF); result.Error != nil {

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

		tmp_citizen := domain.Citizen{ID: element.ID, Name: element.Name, CPF: element.CPF, Birthdate: element.Birthdate}
		citizens = append(citizens, tmp_citizen)
	}

	return citizens, result.Error
}

func (CitizenRepo *mysqlCitizenRepository) CreateCitizen(citizen domain.Citizen) (res domain.Citizen, err error) {
	citizen_model := models.Citizen{Name: utils.Encrypt(citizen.Name),
		CPF: utils.Encrypt(citizen.CPF), Birthdate: citizen.Birthdate}

	if result := CitizenRepo.Conn.Create(&citizen_model); result.Error != nil {

		return citizen, result.Error
	}

	return_citizen := domain.Citizen{ID: citizen_model.ID, Name: citizen_model.Name, CPF: citizen_model.CPF, Birthdate: citizen_model.Birthdate}
	return return_citizen, err
}

// func (CitizenRepo *mysqlCitizenRepository) UpdateCitizenByID(ID int) (res domain.Citizen, err error) {
// 	return
// }

// func (CitizenRepo *mysqlCitizenRepository) InsertNewDebt(ID int) (res domain.Citizen, err error) {
// 	return
// }

// func (CitizenRepo *mysqlCitizenRepository) DeleteDebt(ID int) (res domain.Citizen, err error) {
// 	return
// }

// func (CitizenRepo *mysqlCitizenRepository) UpdateDebt(ID int) (res domain.Citizen, err error) {
// 	return
// }

// func (CitizenRepo *mysqlCitizenRepository) InsertNewAddress(ID int) (res domain.Citizen, err error) {
// 	return
// }

// func (CitizenRepo *mysqlCitizenRepository) DeleteAddress(ID int) (res domain.Citizen, err error) {
// 	return
// }

// func (CitizenRepo *mysqlCitizenRepository) UpdateAddress(ID int) (res domain.Citizen, err error) {
// 	return
// }
