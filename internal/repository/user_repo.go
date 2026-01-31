package repository

import (
	"context"
	"log"
	"main/internal/dao"
	"main/internal/database"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, email, password, firstName, lastName string) (*dao.User, error)
	Get(ctx context.Context, email, password string) (*dao.User, error)
	Update(*dao.User) error
	Delete(*dao.User) error
}

type postgresUserRepository struct {
	db *gorm.DB
}

func NewUserRepo() UserRepository {
	if err := database.DBInstance.AutoMigrate(&dao.User{}); err != nil {
		log.Fatalf("Failed to migrate user schema: %v", err)
	}

	return &postgresUserRepository{
		db: database.DBInstance,
	}
}

func (r *postgresUserRepository) Create(ctx context.Context, email, password, firstName, lastName string) (*dao.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := dao.User{
		Email:     email,
		Password:  string(hashedPassword),
		FirstName: firstName,
		LastName:  lastName,
	}

	if err := r.db.WithContext(ctx).Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *postgresUserRepository) Get(ctx context.Context, email, password string) (*dao.User, error) {
	user := dao.User{}
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	// check password matches
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *postgresUserRepository) Update(user *dao.User) error {
	return r.db.Save(user).Error
}

func (r *postgresUserRepository) Delete(user *dao.User) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// Soft delete tokens first (or after — either is fine)
		if err := tx.Where("user_id = ?", user.ID).Delete(&dao.RefreshToken{}).Error; err != nil {
			return err
		}

		// Soft delete user
		if err := tx.Delete(user).Error; err != nil {
			return err
		}

		return nil
	})
}
