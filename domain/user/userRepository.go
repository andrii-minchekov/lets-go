package usr

type UserRepository interface {
	CreateUser(user User) (int64, error)
	GetUserByEmail(email string) (*User, error)
}
