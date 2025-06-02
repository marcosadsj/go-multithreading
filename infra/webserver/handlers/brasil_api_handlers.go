package handlers

import (
	"encoding/json"
	"fmt"
	"go-multithreading/dto"
	"net/http"
)

type BrasilAPI struct {
	BaseURL string
	Chan    chan<- dto.Address
}

func NewBrasilAPI(baseURL string, ch chan<- dto.Address) *BrasilAPI {
	return &BrasilAPI{
		BaseURL: baseURL,
		Chan:    ch,
	}
}

func (b *BrasilAPI) GetAddressByCep(cep string) (address dto.Address, err error) {
	requestURL := b.BaseURL + "/" + cep

	response, err := http.Get(requestURL)

	if err != nil {
		return dto.Address{}, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return dto.Address{}, fmt.Errorf("error fetching data: %s", response.Status)
	}

	var brasilAPIResponse dto.BrasilAPIResponse

	if err := json.NewDecoder(response.Body).Decode(&brasilAPIResponse); err != nil {
		return dto.Address{}, err
	}

	if err := brasilAPIResponse.Validate(); err != nil {
		return dto.Address{}, err
	}

	address = brasilAPIResponse.ToAddress()

	b.Chan <- address

	return address, nil
}
