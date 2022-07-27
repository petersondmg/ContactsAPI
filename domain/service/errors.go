package service

import "errors"

var (
	ErrInvalidPhone = errors.New("invalid phone")
	ErrInvalidName  = errors.New("invalid name")
)
