package uc

import (
	"errors"
	usr "github.com/andrii-minchekov/lets-go/domain/user"
	"github.com/stretchr/testify/mock"
	assert "github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
	"log"
	"reflect"
	"runtime"
	"testing"
)

func TestNewUserUseCase(t *testing.T) {
	type args struct {
		repo usr.UserRepository
	}
	repo := &mockDbRepo{}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{"should create new instance of struct initialized",
			args{repo},
			userUseCaseImpl{repo, bcrypt.CompareHashAndPassword},
		},
		{"should panic when passed repo param is nil",
			args{nil},
			errRepoIsNil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := recoverPanicIfNeeded(NewUserUseCase, tt.args.repo)
			if err != nil {
				if err != tt.want {
					t.Errorf("NewUserUseCase() panic with err = %v, wantErr %v", err, tt.want)
					return
				}
			} else {
				assert.New(t).EqualValuesf(tt.want.(userUseCaseImpl).Repo, got.(userUseCaseImpl).Repo, "NewUserUseCase func created struct with not expected repo field")
				assert.New(t).EqualValuesf(funcName(tt.want.(userUseCaseImpl).hashComparator), funcName(got.(userUseCaseImpl).hashComparator), "NewUserUseCase func created struct with not expected rehashComparator field")
			}
		})
	}
}

func funcName(fun interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(fun).Pointer()).Name()
}

func recoverPanicIfNeeded(fn func(repo usr.UserRepository) UserUseCase, arg usr.UserRepository) (value interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errRepoIsNil
			log.Printf("recovered panic with expectedErr: %v", err)
		}
	}()
	value = fn(arg)
	return
}

func TestUserUseCase_SignInUser(t *testing.T) {
	type fields struct {
		repo usr.UserRepository
	}
	type args struct {
		email    string
		password string
	}
	expectedUserId := 1
	mockedRepo := func(user *usr.User, err error) usr.UserRepository {
		mockDbRepo := &mockDbRepo{}
		mockDbRepo.On("GetUserByEmail", mock.Anything).Return(user, err)
		return mockDbRepo
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantValue int
		wantErr   error
	}{
		{
			name:      "should create user successfully",
			fields:    fields{mockedRepo(&usr.User{Id: expectedUserId, Password: "password"}, nil)},
			args:      args{"email", "password"},
			wantValue: expectedUserId,
			wantErr:   nil,
		},
		{
			name:      "should return ErrUserAlreadyExist when repo return ErrUserAlreadyExist",
			fields:    fields{mockedRepo(nil, usr.ErrUserAlreadyExist)},
			args:      args{"email", "password"},
			wantValue: 0,
			wantErr:   usr.ErrUserAlreadyExist,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := userUseCaseImpl{
				Repo: tt.fields.repo,
				hashComparator: func(hash []byte, text []byte) error {
					return nil
				},
			}
			got, err := uc.SignInUser(tt.args.email, tt.args.password)
			if (err != nil) && !errors.Is(err, tt.wantErr) {
				t.Errorf("SignInUser() expectedErr = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.wantValue {
				t.Errorf("SignInUser() got = %v, wantValue %v", got, tt.wantValue)
			}
		})
	}
}

type mockDbRepo struct {
	mock.Mock
}

func (r *mockDbRepo) CreateUser(user usr.User) (int, error) {
	panic("implement me")
}

func (r *mockDbRepo) GetUserByEmail(email string) (*usr.User, error) {
	args := r.Called(email)
	return args.Get(0).(*usr.User), args.Error(1)
}
