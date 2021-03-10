package usr

type UserRepository interface {
	CreateUser(user User) (int, error)
	GetUserByEmail(email string) (*User, error)
}
