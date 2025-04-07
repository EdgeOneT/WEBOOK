package service

import (
	"WEBOOK/internal/domain"
	"WEBOOK/internal/respository"
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrDuplicateEmail        = respository.ErrDuplicateEmail
	ErrInvalidUserOrPassword = errors.New("用户不存在或者密码不对")
)

type UserService struct {
	repo *respository.UserRepository
}

func NewUserService(repo *respository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (svc *UserService) Signup(ctx context.Context, u domain.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	return svc.repo.Create(ctx, u)
}
