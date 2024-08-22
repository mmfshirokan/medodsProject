package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mmfshirokan/medodsProject/internal/model"
	"github.com/mmfshirokan/medodsProject/internal/service"
)

type Postgres struct {
	conn *pgxpool.Pool
}

func New(conn *pgxpool.Pool) *Postgres {
	return &Postgres{conn: conn}
}

func (p *Postgres) Add(ctx context.Context, rft model.RefreshToken) error {
	_, err := p.conn.Exec(ctx, "INSERT INTO medods.refresh_token(id, user_id, hash, expiration) VALUES($1, $2, $3, $4)", rft.ID, rft.UserID, rft.Hash, rft.Expiration)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) GetWithUserID(ctx context.Context, id uuid.UUID) ([]model.RefreshToken, error) {
	var res []model.RefreshToken

	rows, err := p.conn.Query(ctx, "SELECT (id, user_id, hash, expiration) FROM medods.refresh_token WHERE user_id = $1", id)
	if err != nil {
		return nil, err
	}

	var rft model.RefreshToken

	for rows.Next() {
		err = rows.Scan(&rft)
		if err != nil {
			return nil, err
		}

		res = append(res, rft)
	}

	return res, nil
}

func (p *Postgres) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := p.conn.Exec(ctx, "DELETE FROM medods.refresh_token WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) GetHash(ctx context.Context, rftID uuid.UUID) (string, error) {
	var (
		uHash string
		exp   time.Time
	)

	err := p.conn.QueryRow(ctx, "SELECT hash, expiration FROM medods.refresh_token WHERE id = $1", rftID).Scan(&uHash, &exp)
	if err != nil {
		return "", err
	}

	if exp.Before(time.Now().UTC()) {
		return "", service.ErrRftTokeExpired
	}

	return uHash, nil
}
