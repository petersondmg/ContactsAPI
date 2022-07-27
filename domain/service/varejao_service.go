package service

import (
	"context"
	"capi/domain/entity"
)

func (s *Service) Varejao() ClientService {
	return &varejaoService{
		mainService: s,
	}
}

type varejaoService struct {
	mainService *Service
}

func (ms *varejaoService) AddContacts(ctx context.Context, contacts []*entity.Contact) error {
	for _, c := range contacts {
		if c.Name == "" {
			return ErrInvalidName
		}

		if !phoneRE.MatchString(c.Phone) {
			return ErrInvalidPhone
		}
	}

	return ms.mainService.varejaoContactRepo.Add(ctx, contacts)
}
