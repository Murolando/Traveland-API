package service

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"strconv"
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
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (s AuthService) CreateUser(user ent.User) (map[string]interface{}, error) {
	realPass := user.Password
	user.Password = s.generateHashPassword(user.Password)
	id, err := s.repo.CreateUser(user, realPass)
	if err != nil {
		return nil, err
	}
	token, err := s.GenerateToken(id)
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{"token": token, "user-id": id}, nil
}
func (s AuthService) SignIn(mail string, password string) (int, error) {
	password = s.generateHashPassword(password)
	id, err := s.repo.GetUserByMailAndPassword(mail, password)
	if err != nil {
		return 0, err
	}
	return id, nil

}

// Для генерации  хэша пароля
func (s AuthService) generateHashPassword(password string) string {

	hash := sha256.Sum256([]byte(password))
	return fmt.Sprintf("%x", hash)
}

// Для генерации jwt токена
func (s AuthService) GenerateToken(id int) (string, error) {
	fmt.Println(id)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
			Subject:   strconv.Itoa(id),
		},
	})
	str, err := token.SignedString([]byte(signInKey))
	return str, err
}

// Распарсивает jwt токен, для проверки его валидности
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
	fmt.Println(124)
	if claims, ok := token.Claims.(*tokenClaims); ok && token.Valid {
		id,err:=strconv.Atoi(claims.Subject)
		if err!=nil{
			return 0, errors.New("Troubles with convert string to int")
		}
		return id, nil
	}
	return 0, errors.New("token claims are not of type *tokenClaims")
}
