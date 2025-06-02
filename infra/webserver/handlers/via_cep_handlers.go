package handlers

import (
	"encoding/json"
	"go-multithreading/dto"
	"net/http"
)

type ViaCep struct {
	BaseURL string
	Chan    chan<- dto.Address
}

func NewViaCep(baseURL string, ch chan<- dto.Address) *ViaCep {
	return &ViaCep{
		BaseURL: baseURL,
		Chan:    ch,
	}
}

func (v *ViaCep) GetAddressByCep(cep string) (address dto.Address, err error) {

	requestURL := v.BaseURL + cep + "/json/"

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
