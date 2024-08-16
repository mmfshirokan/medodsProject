package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	//"github.com/mmfshirokan/medodsProject/internal/model"
)

type Postgres struct {
	conn *pgxpool.Pool
}

func New(conn *pgxpool.Pool) *Postgres {
	return &Postgres{conn: conn}
}

func (p *Postgres) Add(ctx context.Context, rft string) error {
	// TODO implement

	return nil
}

func (p *Postgres) GetWithUserID(ctx context.Context, id uuid.UUID) ([]string, error) {
	// TDOO implement

	return []string{}, nil
}

func (p *Postgres) Delete(ctx context.Context, id uuid.UUID) error {
	// TODO impliment

	return nil
}
