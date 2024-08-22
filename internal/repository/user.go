package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/mmfshirokan/medodsProject/internal/model"
)

// For Testing purposes
func (p *Postgres) AddUsr(ctx context.Context, usr model.User) error {
	_, err := p.conn.Exec(ctx, "INSERT INTO medods.user (id, ip, name, email, password) VALUES ($1, $2, $3, $4, $5)", usr.ID, usr.IP, usr.Name, usr.Email, usr.Password)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) GetPwd(ctx context.Context, uid uuid.UUID) (string, error) {
	var pwd string
	err := p.conn.QueryRow(ctx, "SELECT password FROM medods.user WHERE id = $1", uid).Scan(&pwd)
	if err != nil {
		return "", err
	}

	return pwd, nil
}
