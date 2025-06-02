package handlers

import (
	"encoding/json"
	"fmt"
	"go-multithreading/dto"
	"net/http"
)

type BrasilAPIHandler struct {
	BaseURL string
	Chan    chan<- dto.Address
}

func (b *BrasilAPIHandler) FullURLWithCep(cep string) string {

	return b.BaseURL + "/" + cep
}

func NewBrasilAPIHandler(baseURL string, ch chan<- dto.Address) *BrasilAPIHandler {
	return &BrasilAPIHandler{
		BaseURL: baseURL,
		Chan:    ch,
	}
}

func (b *BrasilAPIHandler) GetAddressByCep(cep string) (address dto.Address, err error) {
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
