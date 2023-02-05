package service

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
	"traveland/ent"
	"traveland/pkg/repository"

	"github.com/dgrijalva/jwt-go"
)

const (
	refreshTokenTime = 30 * 24
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
	user.Password = s.generateHashPassword(user.Password)

	// ses,err := s.NewRefreshToken()
	// if err!=nil{
	// 	return nil, err
	// }

	// user.Session.RefreshToken = ses
	// user.Session.ExpiredAt = time.Now().Add(refreshTokenTime*time.Hour).Unix()

	id, err := s.repo.CreateUser(user)
	if err != nil {
		return nil, err
	}
	token, err := s.GenerateToken(id)
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{"token": token}, nil
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
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(15 * time.Minute).Unix(),
			IssuedAt:  time.Now().Unix(),
			Subject:   strconv.Itoa(id),
		},
	})
	str, err := token.SignedString([]byte(os.Getenv("SIGNINKEY")))
	return str, err
}

// Распарсивает jwt токен, для проверки его валидности
func (s AuthService) ParseToken(accesstoken string) (int, error) {
	token, err := jwt.ParseWithClaims(accesstoken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid string method")
		}
		return []byte(os.Getenv("SIGNINKEY")), nil
	})
	if err != nil {
		return 0, err
	}
	if claims, ok := token.Claims.(*tokenClaims); ok && token.Valid {
		id, err := strconv.Atoi(claims.Subject)
		if err != nil {
			return 0, errors.New("Troubles with convert string to int")
		}
		return id, nil
	}
	return 0, errors.New("token claims are not of type *tokenClaims")
}

// refresh token
func (s AuthService) NewRefreshToken() (string, error) {
	b := make([]byte, 32)
	ss := rand.NewSource(time.Now().Unix())
	r := rand.New(ss)

	_, err := r.Read(b)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", b), nil
}
