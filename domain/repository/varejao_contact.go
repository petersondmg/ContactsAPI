package repository

import (
	"context"
	"database/sql"
	"fmt"
	"capi/domain/entity"
)

func NewVarejaoContact(db *sql.DB) Contact {
	return &varejaoContact{db: db}
}

type varejaoContact struct {
	db *sql.DB
}

func (c *varejaoContact) Add(ctx context.Context, contacts []*entity.Contact) error {
	tx, err := c.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("error starting transaction: %w", err)
	}

	stmt, err := tx.Prepare("INSERT INTO contacts (nome, celular) VALUES ($1, $2) RETURNING id")
	if err != nil {
		return fmt.Errorf("error creating insert stmt: %w", err)
	}

	var success = false

	defer func() {
		if !success {
			tx.Rollback()
		}
	}()

	for _, contact := range contacts {
		var contactID int

		err := stmt.QueryRow(
			contact.Name,
			contact.Phone,
		).Scan(&contactID)
		if err != nil {
			return fmt.Errorf("error inserting contact: %w", err)
		}

		contact.ID = contactID
	}

	success = true

	return tx.Commit()
}
