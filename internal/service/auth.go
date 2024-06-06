package service

import (
	"context"
	"crypto/sha256"
	"fmt"
	"time"

	"rutube-task/internal/config"
	"rutube-task/internal/entity"
	"rutube-task/internal/repository"

	"github.com/golang-jwt/jwt/v4"
)

const (
	// Для теста sender, чтоб каждый раз не получать новый токен, в продакшане так не делаем, не безопасно
	tokenExp = time.Hour * 1
)

// claims - стуктура в котором хранится jwt токен и id пользователя
type claims struct {
	jwt.RegisteredClaims
	UserID int
}

type AuthService struct {
	rep *repository.Repository
	cfg *config.Config
}

func NewAuthService(cfg *config.Config, rep *repository.Repository) *AuthService {
	return &AuthService{
		rep: rep,
		cfg: cfg,
	}
}

func (s *AuthService) CreateUser(ctx context.Context, user entity.User) (int, error) {
	user.Password = s.GeneratePasswordHash(user.Password)
	id, err := s.rep.AuthorizationRepository.SetUserDB(ctx, user)
	if err != nil {
		return 0, fmt.Errorf("CreateUser: %w", err)
	}
	return id, nil
}

// ValidateLogin - метод, который проверяет наличие логина в репозитории
func (s *AuthService) ValidateLogin(ctx context.Context, user entity.User) error {
	err := s.rep.AuthorizationRepository.Validate(ctx, user.Login)

	if err != nil {
		return entity.ErrNotUniqueLogin
	}

	return nil
}

// CheckData - метод, который проверяет пользователя с такими данными
func (s *AuthService) CheckData(ctx context.Context, user entity.User) (int, error) {
	user.Password = s.GeneratePasswordHash(user.Password)
	id, err := s.rep.AuthorizationRepository.GetUserFromDB(ctx, user)
	if err != nil {
		return 0, fmt.Errorf("CheckData: %w", err)
	}

	return id, nil
}

// GenerateJWTToken - создаем токен на основе id пользователя
func (s *AuthService) GenerateJWTToken(userID int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenExp)),
		},
		UserID: userID,
	})

	tokenString, err := token.SignedString([]byte(s.cfg.SecretKey))
	if err != nil {
		return "", fmt.Errorf("GenerateJWTToken: err to generate token: %w", err)
	}

	return tokenString, nil
}

// GetUserIDFromToken - получаем id пользователя по введенному токену
func (s *AuthService) GetUserIDFromToken(tokenString string) int {
	claims := &claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("GetUserIDFromToken: unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(s.cfg.SecretKey), nil
	})

	if err != nil {
		return -1
	}

	if !token.Valid {
		return -1
	}

	return claims.UserID
}

// GeneratePasswordHash - шифрует пароль
func (s *AuthService) GeneratePasswordHash(pass string) string {
	h := sha256.New()
	h.Write([]byte(pass))
	return fmt.Sprintf("%x", h.Sum([]byte(s.cfg.Salt)))

}
