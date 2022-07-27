package service

import (
	"context"
	"regexp"

	"capi/domain/entity"
	"capi/domain/repository"
)

var (
	phoneRE = regexp.MustCompile(`^(55)(\d{2})(\d{4,5})(\d{4})$`)
)

type Service struct {
	macapaContactRepo  repository.Contact
	varejaoContactRepo repository.Contact
}

func New(macapaContactRepo repository.Contact, varejaoContactRepo repository.Contact) *Service {
	return &Service{
		macapaContactRepo:  macapaContactRepo,
		varejaoContactRepo: varejaoContactRepo,
	}
}

type ClientService interface {
	AddContacts(ctx context.Context, contacts []*entity.Contact) error
}
