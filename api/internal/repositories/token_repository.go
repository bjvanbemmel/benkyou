package repositories

import (
	"context"

	"github.com/bjvanbemmel/benkyou/internal/data"
	"github.com/bjvanbemmel/benkyou/internal/database"
	"github.com/bjvanbemmel/benkyou/internal/errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TokenRepository struct {
	conn    *pgxpool.Pool
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

func (t TokenRepository) Close() {
	t.conn.Close()
}

func (t TokenRepository) Index() ([]data.Token, error) {
	tokens, err := t.queries.ListTokens(t.ctx)
	return tokens, errors.MapDatabaseError(err)
}

func (t TokenRepository) Get(id uuid.UUID) (data.Token, error) {
	token, err := t.queries.GetToken(t.ctx, id)
	return token, errors.MapDatabaseError(err)
}

func (t TokenRepository) GetByValue(value string) (data.Token, error) {
	token, err := t.queries.GetTokenByValue(t.ctx, value)
	return token, errors.MapDatabaseError(err)
}

func (t TokenRepository) GetByUserID(id uuid.UUID) (data.Token, error) {
	token, err := t.queries.GetTokenByUserID(t.ctx, id)
	return token, errors.MapDatabaseError(err)
}

func (t TokenRepository) Create(token data.CreateTokenParams) (data.Token, error) {
	created, err := t.queries.CreateToken(t.ctx, token)
	return created, errors.MapDatabaseError(err)
}

func (t TokenRepository) Update(token data.UpdateTokenParams) (data.Token, error) {
	updated, err := t.queries.UpdateToken(t.ctx, token)
	return updated, errors.MapDatabaseError(err)
}

func (t TokenRepository) Delete(id uuid.UUID) error {
	return errors.MapDatabaseError(t.queries.DeleteToken(t.ctx, id))
}
