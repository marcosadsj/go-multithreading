package dto

import (
	"testing"

	default_errors "go-multithreading/infra/errors"
)

func TestBrasilApiValidate(t *testing.T) {
	tests := []struct {
		name              string
		brazilApiResponse BrasilAPIResponse
		expected          error
	}{
		{
			"Valid CEP",
			BrasilAPIResponse{
				Cep:          "86811-190",
				State:        "PR",
				City:         "Apucarana",
				Neighborhood: "Jardim Laranjeiras",
				Street:       "Rua Diniz José Silvério",
				Service:      "open-cep",
			},
			nil,
		},
		{
			"Invalid CEP - Too Short",
			BrasilAPIResponse{
				Cep:          "12345",
				Street:       "Rua Teste",
				Neighborhood: "Bairro Teste",
				City:         "Cidade Teste",
				State:        "TT",
				Service:      "open-cep",
			},
			default_errors.ErrInvalidCepLength,
		},
		{
			"Invalid CEP - Too Long",
			BrasilAPIResponse{
				Cep:          "1234567890",
				Street:       "Rua Teste",
				Neighborhood: "Bairro Teste",
				City:         "Cidade Teste",
				State:        "TT",
				Service:      "open-cep",
			},
			default_errors.ErrInvalidCepLength,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.brazilApiResponse.Validate()
			if err != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, err)
			}
		})
	}
}
