package usecase

import (
	"github.com/Andre-Sacilotti/golang-credit-backend/citizen_api/domain"
)

type CitizenUsecase struct {
	CitizenRepo domain.CitizenRepository
}

func UsecaseInterface(a domain.CitizenRepository) domain.CitizenUsecase {
	return &CitizenUsecase{a}
}

func (CitizenUC *CitizenUsecase) GetCitizenByID(ID int) (res []domain.Citizen) {
	res_, err := CitizenUC.CitizenRepo.GetCitizenByID(ID)
	ArrayRes := []domain.Citizen{res_}

	if err != nil {
		return ArrayRes
	}
	return ArrayRes
}

func (CitizenUC *CitizenUsecase) GetCitizenByCPF(CPF string) (res []domain.Citizen) {
	res_, err := CitizenUC.CitizenRepo.GetCitizenByCPF(CPF)
	ArrayRes := []domain.Citizen{res_}
	if err != nil {
		return ArrayRes
	}
	return ArrayRes
}

func (CitizenUC *CitizenUsecase) CreateCitizen(Citizen domain.Citizen) (res []domain.Citizen, err error) {
	res_, err := CitizenUC.CitizenRepo.CreateCitizen(Citizen)
	ArrayRes := []domain.Citizen{res_}
	if err != nil {
		return ArrayRes, domain.ErrAlreadyExists
	}
	return ArrayRes, err
}

func (CitizenUC *CitizenUsecase) GetDebtsByCitizenId(ID int) (res []domain.Debt) {
	res_, err := CitizenUC.CitizenRepo.GetDebtsByCitizenId(ID)
	if err != nil {
		return res_
	}
	return res_
}

func (CitizenUC *CitizenUsecase) GetAddressByCitizenId(ID int) (res []domain.Address) {
	res_, err := CitizenUC.CitizenRepo.GetAddressByCitizenId(ID)
	if err != nil {
		return res_
	}
	return res_
}

func (CitizenUC *CitizenUsecase) GetAllCitizen() (res []domain.Citizen) {
	res_, err := CitizenUC.CitizenRepo.GetAllCitizen()
	if err != nil {
		return res_
	}
	return res_
}

func (CitizenUC *CitizenUsecase) InsertNewAddress(Address domain.Address, ID int) (res []domain.Address) {
	res_, err := CitizenUC.CitizenRepo.InsertNewAddress(Address, ID)
	ArrayRes := []domain.Address{res_}
	if err != nil {
		return ArrayRes
	}
	return ArrayRes
}

func (CitizenUC *CitizenUsecase) InsertNewDebt(Debt domain.Debt, ID int) (res []domain.Debt) {
	res_, err := CitizenUC.CitizenRepo.InsertNewDebt(Debt, ID)
	ArrayRes := []domain.Debt{res_}
	if err != nil {
		return ArrayRes
	}
	return ArrayRes
}

func (CitizenUC *CitizenUsecase) UpdateCitizenByID(Citizen domain.Citizen, ID int) (res []domain.Citizen) {
	res_, err := CitizenUC.CitizenRepo.UpdateCitizenByID(Citizen, ID)
	ArrayRes := []domain.Citizen{res_}
	if err != nil {
		return ArrayRes
	}
	return ArrayRes
}

func (CitizenUC *CitizenUsecase) DeleteDebt(ID int) (res []domain.Debt) {
	res_, err := CitizenUC.CitizenRepo.DeleteDebt(ID)
	ArrayRes := []domain.Debt{res_}
	if err != nil {
		return ArrayRes
	}
	return ArrayRes
}

func (CitizenUC *CitizenUsecase) DeleteAddress(ID int) (res []domain.Address) {
	res_, err := CitizenUC.CitizenRepo.DeleteAddress(ID)
	ArrayRes := []domain.Address{res_}
	if err != nil {
		return ArrayRes
	}
	return ArrayRes
}

func (CitizenUC *CitizenUsecase) UpdateAddress(Address domain.Address, ID int) (res []domain.Address) {
	res_, err := CitizenUC.CitizenRepo.UpdateAddress(Address, ID)
	ArrayRes := []domain.Address{res_}
	if err != nil {
		return ArrayRes
	}
	return ArrayRes
}

func (CitizenUC *CitizenUsecase) UpdateDebt(Debt domain.Debt, ID int) (res []domain.Debt) {
	res_, err := CitizenUC.CitizenRepo.UpdateDebt(Debt, ID)
	ArrayRes := []domain.Debt{res_}
	if err != nil {
		return ArrayRes
	}
	return ArrayRes
}
