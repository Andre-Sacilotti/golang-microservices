package domain

import "errors"

var (
	ErrInternalServerError = errors.New("Internal Server Error")
	ErrNotAutorized        = errors.New("Not Authorized")
	ErrNotFound            = errors.New("Not Found")
	ErrAlreadyExists       = errors.New("Already Exists")
	ErrCantUpdateCPF       = errors.New("Cant Update CPF")
	ErrCantUpdate          = errors.New("Cant Update Debt")
)
