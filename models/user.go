package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID       int    `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
	Email    string `db:"email"`
}

func (mod User) TableName() string {
	return "users"
}

func (m *User) HashPassword() error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(m.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	m.Password = string(hashed)
	return nil
}

func (m User) VerifyPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(m.Password), []byte(password))
}
