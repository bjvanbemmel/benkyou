package repositories

import (
	"context"

	"github.com/bjvanbemmel/benkyou/internal/data"
	"github.com/bjvanbemmel/benkyou/internal/database"
	"github.com/bjvanbemmel/benkyou/internal/errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ActionPointRepository struct {
	conn    *pgxpool.Pool
	ctx     context.Context
	queries *data.Queries
}

func NewActionPointRepository(ctx context.Context) (ActionPointRepository, error) {
	conn, err := database.New(context.Background())
	if err != nil {
		return ActionPointRepository{}, err
	}

	return ActionPointRepository{
		conn:    conn,
		ctx:     ctx,
		queries: data.New(conn),
	}, nil
}

func (a ActionPointRepository) Index() ([]data.ActionPoint, error) {
	actionPoints, err := a.queries.ListActionPoints(a.ctx)
	return actionPoints, errors.MapDatabaseError(err)
}

func (a ActionPointRepository) Get(id uuid.UUID) (data.ActionPoint, error) {
	actionPoint, err := a.queries.GetActionPoint(a.ctx, id)
	return actionPoint, errors.MapDatabaseError(err)
}

func (a ActionPointRepository) Create(title string) (data.ActionPoint, error) {
	created, err := a.queries.CreateActionPoint(a.ctx, title)
	return created, errors.MapDatabaseError(err)
}

func (a ActionPointRepository) Update(actionPoint data.UpdateActionPointParams) (data.ActionPoint, error) {
	updated, err := a.queries.UpdateActionPoint(a.ctx, actionPoint)
	return updated, errors.MapDatabaseError(err)
}

func (a ActionPointRepository) Delete(id uuid.UUID) error {
	return errors.MapDatabaseError(a.queries.DeleteActionPoint(a.ctx, id))
}
