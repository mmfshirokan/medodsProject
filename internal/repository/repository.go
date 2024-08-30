package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/mmfshirokan/medodsProject/internal/model"
	"gorm.io/gorm"
)

type Postgres struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Postgres {
	return &Postgres{
		db: db,
	}
}

func (p *Postgres) Migrate(ctx context.Context) error {
	err := p.db.AutoMigrate(&model.RefreshToken{}, &model.User{})
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) Add(ctx context.Context, token model.RefreshToken) error {
	err := p.db.Create(&token).Error
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) GetWithUserID(ctx context.Context, userID uuid.UUID) ([]model.RefreshToken, error) {
	var tokens []model.RefreshToken
	err := p.db.Find(&tokens, "user_id = ?", userID).Error
	if err != nil {
		return nil, err
	}

	return tokens, nil
}

func (p *Postgres) Delete(ctx context.Context, id uuid.UUID) error {
	err := p.db.Delete(&model.RefreshToken{}, "id = ?", id).Error
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) GetHash(ctx context.Context, rftID uuid.UUID) (string, error) {
	var rft model.RefreshToken
	err := p.db.First(&rft, "id = ?", rftID).Error
	if err != nil {
		return "", err
	}

	return rft.Hash, nil
}

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
