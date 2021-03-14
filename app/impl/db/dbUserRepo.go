package db

import (
	"errors"
	"github.com/andrii-minchekov/lets-go/domain/user"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type DbUserRepository struct {
	Db Database
}

func NewDbUserRepository() usr.UserRepository {
	return DbUserRepository{GetDatabase()}
}

func (r DbUserRepository) CreateUser(user usr.User) (int64, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)

	if err != nil {
		log.Print(err)
		return 0, err
	}

	stmt := "INSERT INTO users (name, email, password, created) VALUES ($1, $2, $3, CURRENT_DATE) RETURNING ID"

	result, err := r.Db.Query(stmt, user.Name, user.Email, string(hashedPassword))
	if err != nil {
		log.Print(err)
		return 0, err
	}

	if err := result.Next(); !err {
		return 0, errors.New("no rows in result")
	}
	var id int64
	err = result.Scan(&id)
	if err != nil {
		return 0, err
	}
	log.Printf("User was created successfully with id %d", id)
	return id, nil
}

func (r DbUserRepository) GetUserByEmail(email string) (*usr.User, error) {
	row := r.Db.QueryRow("SELECT id, name, email, password FROM users WHERE email = $1", email)
	user := usr.User{}
	if err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password); err != nil {
		return nil, err
	}
	return &user, nil
}
