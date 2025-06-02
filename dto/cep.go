package dto

import (
	default_errors "go-multithreading/infra/errors"
	"strings"
)

type CepInput struct {
	Cep string `json:"cep"`
}

func (b CepInput) Validate() error {
	if b.Cep == "" {
		return default_errors.ErrInvalidCep
	}
	if len(b.GetCep()) < 8 || len(b.GetCep()) > 9 {
		return default_errors.ErrInvalidCep
	}

	return nil
}

func (b CepInput) GetCep() string {
	return strings.Replace(b.Cep, "-", "", -1)
}
