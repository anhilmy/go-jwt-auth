package models

type User struct {
	ID       int    `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
	Email    string `db:"email"`
}

func (mod User) TableName() string {
	return "users"
}
