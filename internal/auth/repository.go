package auth

import (
	"context"
	"database/sql"
	"go-jwt-auth/database"
	errors "go-jwt-auth/internal/errors"
	"go-jwt-auth/models"

	dbx "github.com/go-ozzo/ozzo-dbx"
)

type Repository interface {
	// Get Username
	GetByUsername(ctx context.Context, username string) (*models.User, error)
	// Get Id
	GetById(ctx context.Context, id string) (*models.User, error)
	// Get Email
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	// Create User
	Create(ctx context.Context, user *models.User) error
}

type repository struct {
	db *database.Database
}

func NewRepository(db *database.Database) Repository {
	return repository{db}
}

// Create implements Repository.
func (r repository) Create(ctx context.Context, user *models.User) error {
	if err := r.isUnique(ctx, user); err != nil {
		return err
	}
	return r.db.DB().WithContext(ctx).Model(user).Insert()
}

func (r repository) isUnique(ctx context.Context, user *models.User) error {
	var exists models.User

	err := r.db.DB().WithContext(ctx).Select().Where(dbx.Or(dbx.HashExp{"username": user.Username}, dbx.HashExp{"email": user.Email})).One(&exists)

	if err == sql.ErrNoRows {
		return nil
	} else if err != nil {
		return err
	} else {
		return errors.DuplicationError{}
	}

}

func (r repository) GetByUsername(ctx context.Context, username string) (*models.User, error) {
	where := dbx.HashExp{"username": username}
	return r.get(ctx, where)
}

func (r repository) GetById(ctx context.Context, id string) (*models.User, error) {
	where := dbx.HashExp{"id": id}
	return r.get(ctx, where)
}

func (r repository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	where := dbx.HashExp{"email": email}
	return r.get(ctx, where)
}

func (r repository) get(ctx context.Context, where dbx.Expression) (*models.User, error) {
	var res models.User

	err := r.db.DB().WithContext(ctx).Select().Where(where).One(&res)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &res, nil
}
