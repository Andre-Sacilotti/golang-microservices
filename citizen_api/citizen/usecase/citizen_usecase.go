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

func (CitizenUC *CitizenUsecase) GetDebtsByCitizenCPF(CPF string) (res []domain.Debt) {
	res_, err := CitizenUC.CitizenRepo.GetDebtsByCitizenCPF(CPF)
	if err != nil {
		return res_
	}
	return res_
}

func (CitizenUC *CitizenUsecase) GetAddressByCitizenCPF(CPF string) (res []domain.Address) {
	res_, err := CitizenUC.CitizenRepo.GetAddressByCitizenCPF(CPF)
	if err != nil {
		return res_
	}
	return res_
}

func (CitizenUC *CitizenUsecase) GetAllCitizen(Offset int, Limit int) (res []domain.Citizen) {
	res_, err := CitizenUC.CitizenRepo.GetAllCitizen(Offset, Limit)
	if err != nil {
		return res_
	}
	return res_
}

func (CitizenUC *CitizenUsecase) InsertNewAddress(Address domain.Address, CPF string) (res []domain.Address, err error) {
	res_, err := CitizenUC.CitizenRepo.InsertNewAddress(Address, CPF)
	ArrayRes := []domain.Address{res_}
	if err != nil {
		return ArrayRes, err
	}
	return ArrayRes, err
}

func (CitizenUC *CitizenUsecase) InsertNewDebt(Debt domain.Debt, CPF string) (res []domain.Debt, err error) {
	res_, err := CitizenUC.CitizenRepo.InsertNewDebt(Debt, CPF)
	ArrayRes := []domain.Debt{res_}
	if err != nil {
		return ArrayRes, err
	}
	return ArrayRes, err
}

func (CitizenUC *CitizenUsecase) UpdateCitizenByCPF(Citizen domain.Citizen, CPF string) (res []domain.Citizen) {
	res_, err := CitizenUC.CitizenRepo.UpdateCitizenByCPF(Citizen, CPF)
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

func (CitizenUC *CitizenUsecase) UpdateAddress(Address domain.Address, ID int) (res []domain.Address, err error) {
	res_, err := CitizenUC.CitizenRepo.UpdateAddress(Address, ID)
	ArrayRes := []domain.Address{res_}
	if err != nil {
		return ArrayRes, err
	}
	return ArrayRes, err
}

func (CitizenUC *CitizenUsecase) UpdateDebt(Debt domain.Debt, ID int) (res []domain.Debt, err error) {
	res_, err := CitizenUC.CitizenRepo.UpdateDebt(Debt, ID)
	ArrayRes := []domain.Debt{res_}
	if err != nil {
		return ArrayRes, err
	}
	return ArrayRes, err
}
