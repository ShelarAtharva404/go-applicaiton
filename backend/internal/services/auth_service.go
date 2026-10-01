package services

import (
	"errors"

	"github.com/atharva/fleetpulse-backend/internal/auth"
	"github.com/atharva/fleetpulse-backend/internal/models"
	"github.com/atharva/fleetpulse-backend/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repository *repository.UserRepository
}

func NewAuthService(repository *repository.UserRepository) *AuthService {
	return &AuthService{
		repository: repository,
	}
}

func (s *AuthService) Register(
	name string,
	email string,
	password string,
	role string,
) (*models.User, error) {

	if name == "" || email == "" || password == "" {
		return nil, errors.New("name, email and password are required")
	}

	if role == "" {
		role = "DRIVER"
	}

	if role != "ADMIN" && role != "MANAGER" && role != "DRIVER" {
		return nil, errors.New("invalid role")
	}

	// Check if user already exists
	_, err := s.repository.GetByEmail(email)

	if err == nil {
		return nil, errors.New("user already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	user := &models.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
		Role:     role,
	}

	if err := s.repository.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AuthService) Login(
	email string,
	password string,
) (string, *models.User, error) {

	user, err := s.repository.GetByEmail(email)

	if err != nil {
		return "", nil, errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(password),
	)

	if err != nil {
		return "", nil, errors.New("invalid email or password")
	}

	token, err := auth.GenerateToken(user.ID, user.Role)

	if err != nil {
		return "", nil, errors.New("failed to generate token")
	}

	return token, user, nil
}
