package repositories

import (
	"context"

	"github.com/bjvanbemmel/benkyou/internal/data"
	"github.com/bjvanbemmel/benkyou/internal/database"
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

func (u UserRepository) CloseConn() error {
	return u.conn.Conn().Close(u.ctx)
}

func (u UserRepository) Index() ([]data.ListUsersRow, error) {
	return u.queries.ListUsers(u.ctx)
}

func (u UserRepository) Get(id uuid.UUID) (data.GetUserRow, error) {
	return u.queries.GetUser(u.ctx, id)
}

func (u UserRepository) GetWithPassword(id uuid.UUID) (data.User, error) {
	return u.queries.GetUserWithPassword(u.ctx, id)
}

func (u UserRepository) GetUserWithPasswordByUsername(username string) (data.User, error) {
	return u.queries.GetUserWithPasswordByUsername(u.ctx, username)
}

func (u UserRepository) Create(user data.CreateUserParams) (data.CreateUserRow, error) {
	return u.queries.CreateUser(u.ctx, user)
}

func (u UserRepository) Update(user data.UpdateUserParams) (data.UpdateUserRow, error) {
	return u.queries.UpdateUser(u.ctx, user)
}

func (u UserRepository) Delete(id uuid.UUID) error {
	return u.queries.DeleteUser(u.ctx, id)
}
