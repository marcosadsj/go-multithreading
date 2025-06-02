package dto

import (
	"strings"

	default_errors "go-multithreading/infra/errors"
)

// ViaCEPResponse represents the response structure from ViaCEP API.
type ViaCEPResponse struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func (v ViaCEPResponse) Validate() error {
	if v.Cep == "" {
		return default_errors.ErrInvalidCep
	}
	if v.Logradouro == "" || v.Bairro == "" || v.Localidade == "" || v.Uf == "" {
		return default_errors.ErrIncompleteAddress
	}
	if len(v.Cep) < 8 || len(v.Cep) > 9 {
		return default_errors.ErrInvalidCep
	}
	for _, char := range strings.Replace(v.Cep, "-", "", -1) {
		if char < '0' || char > '9' {
			return default_errors.ErrInvalidCep
		}
	}
	if len(v.Uf) != 2 {
		return default_errors.ErrInvalidState
	}
	if len(v.Ibge) != 7 {
		return default_errors.ErrInvalidIbge
	}
	if len(v.Ddd) < 2 || len(v.Ddd) > 3 {
		return default_errors.ErrInvalidDdd
	}
	if len(v.Siafi) < 4 || len(v.Siafi) > 5 {
		return default_errors.ErrInvalidSiafi
	}
	if len(v.Gia) > 0 && len(v.Gia) != 8 {
		return default_errors.ErrInvalidGia
	}
	if len(v.Regiao) == 0 {
		return default_errors.ErrInvalidRegion
	}
	if len(v.Estado) == 0 {
		return default_errors.ErrInvalidState
	}
	return nil
}

func (v ViaCEPResponse) GetCep() string {
	return strings.Replace(v.Cep, "-", "", -1)
}

func (v ViaCEPResponse) ToAddress() Address {
	return Address{
		Cep:          v.GetCep(),
		Street:       v.Logradouro,
		Neighborhood: v.Bairro,
		City:         v.Localidade,
		State:        v.Uf,
	}
}
