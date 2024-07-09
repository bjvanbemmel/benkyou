package repositories

import (
	"context"

	"github.com/bjvanbemmel/benkyou/internal/data"
	"github.com/bjvanbemmel/benkyou/internal/database"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TokenRepository struct {
	conn    *pgxpool.Conn
	ctx     context.Context
	queries *data.Queries
}

func NewTokenRepository(ctx context.Context) (TokenRepository, error) {
	conn, err := database.New(context.Background())
	if err != nil {
		return TokenRepository{}, err
	}

	return TokenRepository{
		conn:    conn,
		ctx:     ctx,
		queries: data.New(conn),
	}, nil
}

func (t TokenRepository) CloseConn() error {
	return t.conn.Conn().Close(t.ctx)
}

func (t TokenRepository) Index() ([]data.Token, error) {
	return t.queries.ListTokens(t.ctx)
}

func (t TokenRepository) Get(id uuid.UUID) (data.Token, error) {
	return t.queries.GetToken(t.ctx, id)
}

func (t TokenRepository) GetByValue(value string) (data.Token, error) {
	return t.queries.GetTokenByValue(t.ctx, value)
}

func (t TokenRepository) GetByUserID(id uuid.UUID) (data.Token, error) {
	return t.queries.GetTokenByUserID(t.ctx, id)
}

func (t TokenRepository) Create(token data.CreateTokenParams) (data.Token, error) {
	return t.queries.CreateToken(t.ctx, token)
}

func (t TokenRepository) Update(token data.UpdateTokenParams) (data.Token, error) {
	return t.queries.UpdateToken(t.ctx, token)
}

func (t TokenRepository) Delete(id uuid.UUID) error {
	return t.queries.DeleteToken(t.ctx, id)
}
