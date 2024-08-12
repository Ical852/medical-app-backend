package services

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"medical-app-backend/models"
	"medical-app-backend/repositories"
	"time"
)

type UserService interface {
	Register(name, email, password string) (*models.User, error)
	Login(email, password string) (string, error)
}

type userService struct {
	userRepository repositories.UserRepository
	jwtSecret      string
}

func NewUserService(repo repositories.UserRepository, secret string) UserService {
	return &userService{userRepository: repo, jwtSecret: secret}
}

func (s *userService) Register(name, email, password string) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
		Role:     "patient", // default role
	}

	err = s.userRepository.Create(user)
	return user, err
}

func (s *userService) Login(email, password string) (string, error) {
	user, err := s.userRepository.FindByEmail(email)
	if err != nil {
		return "", errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(s.jwtSecret))
	return tokenString, err
}
