package postgres

import (
	"context"
	"gazuberlandia"

	_ "github.com/jackc/pgx/v4"
	"github.com/jmoiron/sqlx"
)

type userService struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *userService {
	return &userService{db: db}
}

func (p *userService) CreateUser(ctx context.Context, user *gazuberlandia.User) error {
	tx, err := p.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, "insert into users (name, email, password) values ($1, $2, $3)",
		&user.Name, &user.Email, &user.Password)

	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (p *userService) FindUserById(ctx context.Context, id int) ([]*gazuberlandia.User, error) {
	user := []*gazuberlandia.User{}
	err := p.db.Select(&user, "select * from users u where u.id=$1", id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (p *userService) FindUserByEmail(ctx context.Context, email string) (gazuberlandia.User, error) {
	var user gazuberlandia.User

	err := p.db.GetContext(ctx, &user, "select * from users u where u.email=$1", email)

	if err != nil {
		return gazuberlandia.User{}, err
	}

	return user, nil
}

func (p *userService) FindAllUsers(ctx context.Context) ([]gazuberlandia.User, error) {
	user := []gazuberlandia.User{}

	err := p.db.Select(&user, "select * from users")

	if err != nil {
		return nil, err
	}

	return user, nil
}
