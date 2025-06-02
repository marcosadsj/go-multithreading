package main

import (
	"fmt"
	"go-multithreading/dto"
	"go-multithreading/infra/webserver/handlers"
	"time"
)

func main() {

	brasilApiChan := make(chan dto.Address)
	viaCepChan := make(chan dto.Address)

	brasilApiHandler := handlers.NewBrasilAPIHandler("https://brasilapi.com.br/api/cep/v1", brasilApiChan)
	viaCepHandler := handlers.NewViaCepHandlerHandler("http://viacep.com.br/ws/", viaCepChan)

	cep := dto.CepInput{
		Cep: "86811-190",
	}

	go func() {
		_, err := brasilApiHandler.GetAddressByCep(cep)
		if err != nil {
			fmt.Println("Error fetching address from Brasil API:", err)
			return
		}
	}()

	go func() {
		_, err := viaCepHandler.GetAddressByCep(cep)
		if err != nil {
			fmt.Println("Error fetching address from ViaCEP:", err)
			return
		}
	}()

	select {
	case address1 := <-brasilApiChan:
		fmt.Println("Address from Brasil API: ", address1)
	case address2 := <-viaCepChan:
		fmt.Println("Address from ViaCEP: ", address2)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout: No address received within 1 second")
		return
	}
}
