package service

import (
	"crypto/sha1"
	"fmt"

	project "github.com/Chazarov/rest-app/entities"
	"github.com/Chazarov/rest-app/pkg/repository"
)

const salt = "ask--dfHJdl123!+"

type AuthService struct {
	repo repository.Authorisation
}

func NewAuthService(repo repository.Authorisation) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user project.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
