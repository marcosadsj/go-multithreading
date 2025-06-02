package handlers

import "go-multithreading/dto"

type CepHandler interface {
	// GetAddressByCep retrieves address details based on the provided CEP.
	GetAddressByCep(cep string) (dto.Address, error)
}
