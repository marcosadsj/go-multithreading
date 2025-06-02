package handlers

import (
	"encoding/json"
	"go-multithreading/dto"
	"net/http"
)

type ViaCepHandler struct {
	BaseURL string
	Chan    chan<- dto.Address
}

func (v *ViaCepHandler) FullURLWithCep(cep string) string {
	return v.BaseURL + cep + "/json/"
}

func NewViaCepHandlerHandler(baseURL string, ch chan<- dto.Address) *ViaCepHandler {
	return &ViaCepHandler{
		BaseURL: baseURL,
		Chan:    ch,
	}
}

func (v *ViaCepHandler) GetAddressByCep(cepInput dto.CepInput) (address dto.Address, err error) {

	if err := cepInput.Validate(); err != nil {
		return dto.Address{}, err
	}

	requestURL := v.FullURLWithCep(cepInput.GetCep())

	response, err := http.Get(requestURL)

	if err != nil {
		return dto.Address{}, err
	}

	var viaCepResponse dto.ViaCEPResponse

	if response.StatusCode != http.StatusOK {

		defer response.Body.Close()

		json.NewDecoder(response.Body).Decode(&viaCepResponse)
	}

	if err := viaCepResponse.Validate(); err != nil {
		return dto.Address{}, err
	}

	address = viaCepResponse.ToAddress()

	v.Chan <- address

	return address, nil
}
