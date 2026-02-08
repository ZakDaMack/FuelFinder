package repository

import (
	"context"
	"log"
	"main/internal/dao"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RefreshTokenRepository interface {
	Create(ctx context.Context, userID, tokenID uuid.UUID) (*dao.RefreshToken, error)
	Get(ctx context.Context, tokenID uuid.UUID) (*dao.RefreshToken, error)
	Delete(*dao.RefreshToken) error
}

type postgresRefreshTokenRepository struct {
	db *gorm.DB
}

func NewRefreshTokenRepo(db *gorm.DB) RefreshTokenRepository {
	if err := db.AutoMigrate(&dao.RefreshToken{}); err != nil {
		log.Fatalf("Failed to migrate refresh token schema: %v", err)
	}
	return &postgresRefreshTokenRepository{
		db: db,
	}
}

func (r *postgresRefreshTokenRepository) Create(ctx context.Context, userID, tokenID uuid.UUID) (*dao.RefreshToken, error) {
	refreshToken := dao.RefreshToken{
		UserID:  userID,
		TokenID: tokenID,
	}

	if err := r.db.WithContext(ctx).Create(&refreshToken).Error; err != nil {
		return nil, err
	}

	return &refreshToken, nil
}

func (r *postgresRefreshTokenRepository) Get(ctx context.Context, tokenID uuid.UUID) (*dao.RefreshToken, error) {
	refreshToken := dao.RefreshToken{}
	if err := r.db.WithContext(ctx).Where("token_id = ?", tokenID).First(&refreshToken).Error; err != nil {
		return nil, err
	}

	return &refreshToken, nil
}

func (r *postgresRefreshTokenRepository) Delete(refreshToken *dao.RefreshToken) error {
	return r.db.Delete(refreshToken).Error
}
