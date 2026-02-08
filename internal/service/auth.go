package service

import (
	"context"
	"fmt"
	"log"
	"main/internal/dao"
	"main/internal/repository"
	"main/internal/token"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthService interface {
	Login(ctx context.Context, email, password string) (string, string, error)
	Logout(ctx context.Context, refreshToken string) error
	Register(ctx context.Context, email, password, confirmPassword, firstName, lastName string) (string, string, error)
	RefreshToken(ctx context.Context, refreshToken string) (string, error)
	GetWellKnown(ctx context.Context) (*token.JWKS, error)
}

type authService struct {
	userRepo         repository.UserRepository
	refreshTokenRepo repository.RefreshTokenRepository
	jwtProvider      *token.JWTProvider
	jwks             *token.JWKS
}

const (
	PasswordMismatchError = "passwords do not match"
	InvalidTokenError     = "invalid refresh token"
	UserNotFoundError     = "user not found"
)

func NewAuthService(db *gorm.DB, jwtOptions *token.Options) AuthService {
	jwks, err := token.NewJWKSFromKey(jwtOptions.PrivateKey)
	if err != nil {
		log.Fatalf("failed to generate JWKS: %v", err)
	}
	return &authService{
		userRepo:         repository.NewUserRepo(db),
		refreshTokenRepo: repository.NewRefreshTokenRepo(db),
		jwtProvider:      token.NewJWTProvider(jwtOptions),
		jwks:             jwks,
	}
}

func (s *authService) Login(ctx context.Context, email, password string) (string, string, error) {
	user, err := s.userRepo.Get(ctx, email, password)
	if err != nil {
		return "", "", fmt.Errorf(UserNotFoundError)
	}

	jwtToken, err := s.jwtProvider.CreateJWT(user.Email, user.FirstName, user.LastName)
	if err != nil {
		return "", "", err
	}

	refreshToken, tokenID, err := s.jwtProvider.CreateRefreshJWT()
	if err != nil {
		return "", "", err
	}

	// save refresh token to store
	_, err = s.refreshTokenRepo.Create(ctx, user.ID, tokenID)
	if err != nil {
		return "", "", err
	}

	return jwtToken, refreshToken, nil
}

func (s *authService) Logout(ctx context.Context, refreshToken string) error {
	// check token is valid
	claim, err := s.jwtProvider.ValidateJWT(refreshToken)
	if err != nil {
		return fmt.Errorf(InvalidTokenError)
	}

	// delete refresh token from store if applicable
	s.refreshTokenRepo.Delete(&dao.RefreshToken{TokenID: uuid.MustParse(claim.ID)})
	return nil
}

func (s *authService) Register(ctx context.Context, email, password, confirmPassword, firstName, lastName string) (string, string, error) {
	if password != confirmPassword {
		return "", "", fmt.Errorf(PasswordMismatchError)
	}

	_, err := s.userRepo.Create(ctx, email, password, firstName, lastName)
	if err != nil {
		return "", "", err
	}

	return s.Login(ctx, email, password)
}

func (s *authService) RefreshToken(ctx context.Context, refreshToken string) (string, error) {
	// validate jwt
	claim, err := s.jwtProvider.ValidateJWT(refreshToken)
	if err != nil {
		return "", fmt.Errorf(InvalidTokenError)
	}

	// check token is valid in store
	_, err = s.refreshTokenRepo.Get(ctx, uuid.MustParse(claim.ID))
	if err != nil {
		return "", fmt.Errorf(InvalidTokenError)
	}

	refreshToken, tokenID, err := s.jwtProvider.CreateRefreshJWT()
	if err != nil {
		return "", err
	}

	// save token to store
	_, err = s.refreshTokenRepo.Create(ctx, uuid.MustParse(claim.Subject), tokenID)
	if err != nil {
		return "", err
	}

	return refreshToken, nil
}

func (s *authService) GetWellKnown(ctx context.Context) (*token.JWKS, error) {
	return s.jwks, nil
}
