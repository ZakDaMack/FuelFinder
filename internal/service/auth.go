package service

import (
	"context"
	"fmt"
	"main/internal/dao"
	"main/internal/repository"
	"main/internal/token"

	"github.com/google/uuid"
)

type AuthService interface {
	Login(ctx context.Context, email, password string) (string, string, error)
	Logout(ctx context.Context, refreshToken string) error
	Register(ctx context.Context, email, password, confirmPassword, firstName, lastName string) (string, string, error)
	RefreshToken(ctx context.Context, refreshToken string) (string, error)
}

type authService struct {
	userRepo         repository.UserRepository
	refreshTokenRepo repository.RefreshTokenRepository
	jwtProvider      token.JWTProvider
	jkws             token.JWKS
}

func NewAuthService(jwtOptions token.JWTOptions) AuthService {
	return &authService{
		userRepo:         repository.NewUserRepo(),
		refreshTokenRepo: repository.NewRefreshTokenRepo(),
		jwtProvider:      token.NewJWTProvider(jwtOptions),
		jkws:             token.NewJWKS(jwtOptions.privateKey),
	}
}

func (s *authService) Login(ctx context.Context, email, password string) (string, string, error) {
	user, err := s.userRepo.Get(ctx, email, password)
	if err != nil {
		return "", "", err
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
		return fmt.Errorf("invalid refresh token: %w", err)
	}

	// delete refresh token from store if applicable
	s.refreshTokenRepo.Delete(&dao.RefreshToken{TokenID: uuid.MustParse(claim.ID)})
	return nil
}

func (s *authService) Register(ctx context.Context, email, password, confirmPassword, firstName, lastName string) (string, string, error) {
	if password != confirmPassword {
		return "", "", fmt.Errorf("passwords do not match")
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
		return "", fmt.Errorf("invalid refresh token: %w", err)
	}

	// check token is valid
	_, err = s.refreshTokenRepo.Get(ctx, uuid.MustParse(claim.ID))
	if err != nil {
		return "", fmt.Errorf("refresh token not found: %w", err)
	}

	refreshToken, tokenID, err := s.jwtProvider.CreateRefreshJWT()
	if err != nil {
		return "", err
	}

	// save token to store
	_, err = s.refreshTokenRepo.Create(ctx, user.id, tokenID)
	if err != nil {
		return "", err
	}

	return refreshToken, nil
}
