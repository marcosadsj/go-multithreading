package dto

import (
	"strings"

	default_errors "go-multithreading/infra/errors"
)

// BrasilAPIResponse represents the response structure from BrasilAPI for CEP queries.
type BrasilAPIResponse struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
}

type BrasilAPIInput struct {
	Cep string `json:"cep"`
}

func (b BrasilAPIResponse) Validate() error {
	if b.Cep == "" {
		return default_errors.ErrInvalidCep
	}
	if b.Street == "" || b.Neighborhood == "" || b.City == "" || b.State == "" {
		return default_errors.ErrIncompleteAddress
	}

	if len(b.GetCep()) < 8 || len(b.GetCep()) > 9 {
		return default_errors.ErrInvalidCepLength
	}
	if len(b.State) != 2 {
		return default_errors.ErrInvalidState
	}
	if len(b.Service) == 0 {
		return default_errors.ErrInvalidService
	}
	if len(b.City) == 0 {
		return default_errors.ErrInvalidCity
	}
	if len(b.Neighborhood) == 0 {
		return default_errors.ErrInvalidNeighborhood
	}
	if len(b.Street) == 0 {
		return default_errors.ErrInvalidStreet
	}

	return nil
}

func (b BrasilAPIResponse) GetCep() string {
	return strings.Replace(b.Cep, "-", "", -1)
}

func (b BrasilAPIResponse) ToAddress() Address {
	return Address{
		Cep:          b.GetCep(),
		Street:       b.Street,
		Neighborhood: b.Neighborhood,
		City:         b.City,
		State:        b.State,
	}
}

func (b BrasilAPIInput) Validate() error {
	if b.Cep == "" {
		return default_errors.ErrInvalidCep
	}
	if len(b.Cep) < 8 || len(b.Cep) > 9 {
		return default_errors.ErrInvalidCep
	}
	for _, char := range b.Cep {
		if char < '0' || char > '9' {
			return default_errors.ErrInvalidCep
		}
	}
	return nil
}

func (b BrasilAPIInput) GetCep() string {
	return strings.Replace(b.Cep, "-", "", -1)
}
