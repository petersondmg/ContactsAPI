package service

import (
	"context"
	"fmt"
	"capi/domain/entity"
	"strings"
)

func (s *Service) Macapa() ClientService {
	return &macapaService{
		mainService: s,
	}
}

type macapaService struct {
	mainService *Service
}

func (ms *macapaService) AddContacts(ctx context.Context, contacts []*entity.Contact) error {
	for _, c := range contacts {
		if c.Name == "" {
			return ErrInvalidName
		}

		matches := phoneRE.FindStringSubmatch(c.Phone)
		if len(matches) == 0 {
			return ErrInvalidPhone
		}

		c.Name = strings.ToUpper(c.Name)
		c.Phone = fmt.Sprintf("+%s (%s) %s-%s", matches[1], matches[2], matches[3], matches[4])
	}

	return ms.mainService.macapaContactRepo.Add(ctx, contacts)
}

func (macapaService) validateContact(c *entity.Contact) error {
	if c.Name == "" {
		return ErrInvalidName
	}

	if phoneRE.MatchString(c.Phone) {
		return nil
	}

	return ErrInvalidPhone
}
