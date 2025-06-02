package dto

import (
	"errors"
	default_errors "go-multithreading/infra/errors"
	"testing"
)

func TestCepValidation(t *testing.T) {
	tests := []struct {
		name    string
		cep     CepInput
		wantErr error
	}{
		{
			name:    "Valid CEP",
			cep:     CepInput{Cep: "86811-190"},
			wantErr: nil,
		},
		{
			name:    "Invalid CEP - empty",
			cep:     CepInput{Cep: ""},
			wantErr: default_errors.ErrInvalidCep,
		},
		{
			name:    "Invalid CEP - too short",
			cep:     CepInput{Cep: "86811"},
			wantErr: default_errors.ErrInvalidCep,
		},
		{
			name:    "Invalid CEP - too long",
			cep:     CepInput{Cep: "86811-19000"},
			wantErr: default_errors.ErrInvalidCep,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.cep.Validate()
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
