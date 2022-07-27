package repository

import (
	"context"
	"database/sql"
	"fmt"
	"capi/domain/entity"
)

func NewMacapaContact(db *sql.DB) Contact {
	return &macapaContact{db: db}
}

type Contact interface {
	Add(ctx context.Context, contacts []*entity.Contact) error
}

type macapaContact struct {
	db *sql.DB
}

func (c *macapaContact) Add(ctx context.Context, contacts []*entity.Contact) error {
	tx, err := c.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("error starting transaction: %w", err)
	}

	stmt, err := tx.Prepare("INSERT INTO contacts (nome, celular) VALUES (?, ?)")
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
		res, err := stmt.ExecContext(ctx, contact.Name, contact.Phone)
		if err != nil {
			return fmt.Errorf("error inserting contact: %w", err)
		}

		id, err := res.LastInsertId()
		if err != nil {
			return fmt.Errorf("error getting latest id: %w", err)
		}

		contact.ID = int(id)
	}

	success = true

	return tx.Commit()
}
