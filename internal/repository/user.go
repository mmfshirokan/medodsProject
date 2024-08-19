package repository

import (
	"context"

	"github.com/google/uuid"
)

func (p *Postgres) ValidatePWD(ctx context.Context, id uuid.UUID, pwd string) bool {
	// TODO implement

	return false
}
