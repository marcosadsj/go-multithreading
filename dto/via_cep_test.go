package dto

import (
	"testing"

	default_errors "go-multithreading/infra/errors"
)

func TestViaCepValidate(t *testing.T) {
	tests := []struct {
		name           string
		viaCepResponse ViaCEPResponse
		expected       error
	}{
		{
			"Valid CEP",
			ViaCEPResponse{
				Cep:         "86811-190",
				Logradouro:  "Rua Diniz José Silvério",
				Complemento: "",
				Unidade:     "",
				Bairro:      "Jardim Laranjeiras",
				Localidade:  "Apucarana",
				Uf:          "PR",
				Estado:      "Paraná",
				Regiao:      "Sul",
				Ibge:        "4101408",
				Gia:         "",
				Ddd:         "43",
				Siafi:       "7425",
			},
			nil},
		{
			"Invalid CEP - Too Short",
			ViaCEPResponse{
				Cep:         "12345",
				Logradouro:  "Rua Teste",
				Complemento: "",
				Unidade:     "",
				Bairro:      "Bairro Teste",
				Localidade:  "Cidade Teste",
				Uf:          "TT",
				Estado:      "Estado Teste",
				Regiao:      "Região Teste",
				Ibge:        "0000000",
				Gia:         "",
				Ddd:         "00",
				Siafi:       "0000",
			},
			default_errors.ErrInvalidCep,
		},
		{
			"Invalid CEP - Too Long",
			ViaCEPResponse{
				Cep:         "12345678901234567890",
				Logradouro:  "Rua Teste",
				Complemento: "",
				Unidade:     "",
				Bairro:      "Bairro Teste",
				Localidade:  "Cidade Teste",
				Uf:          "TT",
				Estado:      "Estado Teste",
				Regiao:      "Região Teste",
				Ibge:        "0000000",
				Gia:         "",
				Ddd:         "00",
				Siafi:       "0000",
			},
			default_errors.ErrInvalidCep,
		},
		{
			"Invalid CEP - Non-Numeric",
			ViaCEPResponse{
				Cep:         "ABCDE",
				Logradouro:  "Rua Teste",
				Complemento: "",
				Unidade:     "",
				Bairro:      "Bairro Teste",
				Localidade:  "Cidade Teste",
				Uf:          "TT",
				Estado:      "Estado Teste",
				Regiao:      "Região Teste",
				Ibge:        "0000000",
				Gia:         "",
				Ddd:         "00",
				Siafi:       "0000",
			},
			default_errors.ErrInvalidCep,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.viaCepResponse.Validate()
			if err != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, err)
			}
		})
	}
}
