package repositories

import (
	"context"

	"github.com/bjvanbemmel/benkyou/internal/data"
	"github.com/bjvanbemmel/benkyou/internal/database"
	"github.com/bjvanbemmel/benkyou/internal/errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	conn    *pgxpool.Conn
	ctx     context.Context
	queries *data.Queries
}

func NewUserRepository(ctx context.Context) (UserRepository, error) {
	conn, err := database.New(context.Background())
	if err != nil {
		return UserRepository{}, err
	}

	return UserRepository{
		conn:    conn,
		ctx:     ctx,
		queries: data.New(conn),
	}, nil
}

func (u *UserRepository) CloseConn() error {
	return u.conn.Conn().Close(u.ctx)
}

func (u UserRepository) Index() ([]data.ListUsersRow, error) {
	users, err := u.queries.ListUsers(u.ctx)
	return users, errors.MapDatabaseError(err)
}

func (u UserRepository) Get(id uuid.UUID) (data.GetUserRow, error) {
	user, err := u.queries.GetUser(u.ctx, id)
	return user, errors.MapDatabaseError(err)
}

func (u UserRepository) GetWithPassword(id uuid.UUID) (data.User, error) {
	user, err := u.queries.GetUserWithPassword(u.ctx, id)
	return user, errors.MapDatabaseError(err)
}

func (u UserRepository) GetWithPasswordByEmail(email string) (data.User, error) {
	user, err := u.queries.GetUserWithPasswordByEmail(u.ctx, email)
	return user, err
}

func (u UserRepository) Create(user data.CreateUserParams) (data.CreateUserRow, error) {
	created, err := u.queries.CreateUser(u.ctx, user)
	return created, errors.MapDatabaseError(err)
}

func (u UserRepository) Update(user data.UpdateUserParams) (data.UpdateUserRow, error) {
	updated, err := u.queries.UpdateUser(u.ctx, user)
	return updated, errors.MapDatabaseError(err)
}

func (u UserRepository) Delete(id uuid.UUID) error {
	return errors.MapDatabaseError(u.queries.DeleteUser(u.ctx, id))
}
