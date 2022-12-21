package service

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"time"
	"traveland/ent"
	"traveland/pkg/repository"

	"github.com/dgrijalva/jwt-go"
)

const (
	signInKey = "as8129ru129fijwi9hahg7"
)

type AuthService struct {
	repo repository.Authorization
}

type tokenClaims struct {
	jwt.StandardClaims
	userId int `json:"user_id"`
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (s AuthService) CreateUser(user ent.User) (int, error) {
	user.Password = s.generateHashPassword(user.Password)
	return s.repo.CreateUser(user)
}

// Для генерации  хэша пароля
func (s AuthService) generateHashPassword(password string) string {

	hash := sha256.Sum256([]byte(password))
	return fmt.Sprintf("%x", hash)
}

// Для генерации jwt токена
func (s AuthService) GenerateToken(mail string, password string) (string, error) {
	userId, err := s.repo.GetUserByMailAndPassword(mail, s.generateHashPassword(password))
	if err != nil {
		if userId == -1{
			return "", fmt.Errorf("wrong password")
		}
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(), 
			IssuedAt:  time.Now().Unix(),
		},
		userId: userId,
	})

	return token.SignedString([]byte(signInKey))
}

// Сам хз как это работаети устроено !
func (s AuthService) ParseToken(accesstoken string) (int, error) {
	token, err := jwt.ParseWithClaims(accesstoken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid string method")
		}
		return []byte(signInKey), nil
	})
	if err != nil {
		return 0, err
	}
	if claims, ok := token.Claims.(*tokenClaims); ok && token.Valid {
		return claims.userId, nil
	}
	return 0, errors.New("token claims are not of type *tokenClaims")
}
