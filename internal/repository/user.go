package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/mmfshirokan/medodsProject/internal/model"
)

func (p *Postgres) AddUsr(ctx context.Context, usr model.User) error {
	err := p.db.Create(&usr).Error
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) GetPwd(ctx context.Context, uid uuid.UUID) (string, error) {
	var usr model.User
	err := p.db.First(&usr, "id = ?", uid).Error
	if err != nil {
		return "", err
	}

	return usr.Password, nil
}
