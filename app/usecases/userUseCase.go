package uc

import (
	"database/sql"
	"errors"
	"github.com/andrii-minchekov/lets-go/domain/user"
	"golang.org/x/crypto/bcrypt"
	"log"
)

var errRepoIsNil = errors.New("repo shouldn't be nil")

type UserUseCase interface {
	SignupUser(user usr.User) (int, error)
	SignInUser(email, password string) (int, error)
}

type UserUseCaseImplTest struct {
	Repo usr.Repository
	fun  func()
}

type userUseCaseImpl struct {
	Repo           usr.Repository
	hashComparator hashComparator
}

func NewUserUseCase(repo usr.Repository) UserUseCase {
	if repo == nil {
		log.Panicf(errRepoIsNil.Error())
	}
	return userUseCaseImpl{Repo: repo, hashComparator: bcrypt.CompareHashAndPassword}
}

func (uc userUseCaseImpl) SignupUser(user usr.User) (int, error) {
	return uc.Repo.CreateUser(user)
}

func (uc userUseCaseImpl) SignInUser(email, password string) (int, error) {
	user, err := uc.Repo.GetUserByEmail(email)

	if err == sql.ErrNoRows {
		return 0, usr.ErrInvalidCredentials
	} else if err != nil {
		return 0, err
	}

	err = uc.hashComparator([]byte(user.Password), []byte(password))

	if err == bcrypt.ErrMismatchedHashAndPassword {
		return 0, usr.ErrInvalidCredentials
	} else if err != nil {
		return 0, err
	}

	return user.Id, nil
}

type hashComparator func(hash []byte, text []byte) error
