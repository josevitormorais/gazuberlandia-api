package postgres

import (
	"context"
	"errors"
	"gazuberlandia"

	_ "github.com/jackc/pgx/v4"
	"github.com/jmoiron/sqlx"
)

type database struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *database {
	return &database{db: db}
}

func (p *database) CreateUser(ctx context.Context, user *gazuberlandia.User) error {
	tx := p.db.MustBegin()
	tx.MustExec("insert into users (name, email, password) values ($1, $2, $3)",
		&user.Name, &user.Email, &user.Password)
	return tx.Commit()
}

func (p *database) FindUserById(ctx context.Context, id int) ([]gazuberlandia.User, error) {
	user := []gazuberlandia.User{}
	err := p.db.Select(&user, "select * from users u where u.id=$1", id)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return user, nil
}

func (p *database) FindUserByEmail(ctx context.Context, email string) (gazuberlandia.User, error) {
	var user gazuberlandia.User

	err := p.db.GetContext(ctx, &user, "select * from users u where u.email=$1", email)

	if err != nil {
		return gazuberlandia.User{}, errors.New("user not found")
	}

	return user, nil
}

func (p *database) FindAllUsers(ctx context.Context) ([]gazuberlandia.User, error) {
	user := []gazuberlandia.User{}

	err := p.db.Select(&user, "select * from users")

	if err != nil {
		return nil, errors.New(err.Error())
	}
	return user, nil
}
