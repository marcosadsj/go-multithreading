package dto

import "errors"

var (
	ErrInvalidCep          = errors.New("invalid CEP format")
	ErrCepNotFound         = errors.New("CEP not found")
	ErrIncompleteAddress   = errors.New("incomplete address details")
	ErrInvalidState        = errors.New("invalid state format")
	ErrInvalidIbge         = errors.New("invalid IBGE code format")
	ErrInvalidDdd          = errors.New("invalid DDD format")
	ErrInvalidSiafi        = errors.New("invalid SIAFI code format")
	ErrInvalidGia          = errors.New("invalid GIA code format")
	ErrInvalidRegion       = errors.New("invalid region format")
	ErrInvalidCepFormat    = errors.New("CEP must be 8 digits long")
	ErrInvalidCepLength    = errors.New("CEP must be 8 or 9 characters long")
	ErrInvalidCepCharacter = errors.New("CEP must contain only numeric characters")
	ErrInvalidService      = errors.New("invalid service format")
	ErrInvalidCity         = errors.New("invalid city format")
	ErrInvalidNeighborhood = errors.New("invalid neighborhood format")
	ErrInvalidStreet       = errors.New("invalid street format")
	ErrInvalidAddress      = errors.New("invalid address details")
)
