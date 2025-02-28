package repositories

import (
	"context"

	"github.com/bjvanbemmel/benkyou/internal/data"
	"github.com/bjvanbemmel/benkyou/internal/database"
	"github.com/bjvanbemmel/benkyou/internal/errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type RequirementRepository struct {
	conn    *pgxpool.Conn
	ctx     context.Context
	queries *data.Queries
}

func NewRequirementRepository(ctx context.Context) (RequirementRepository, error) {
	conn, err := database.New(context.Background())
	if err != nil {
		return RequirementRepository{}, err
	}

	return RequirementRepository{
		conn:    conn,
		ctx:     ctx,
		queries: data.New(conn),
	}, nil
}

func (r *RequirementRepository) CloseConn() error {
	return r.conn.Conn().Close(r.ctx)
}

func (r RequirementRepository) Index() ([]data.Sprint, error) {
	sprints, err := r.queries.ListSprints(r.ctx)
	return sprints, errors.MapDatabaseError(err)
}

func (r RequirementRepository) Get(id uuid.UUID) (data.Sprint, error) {
	sprint, err := r.queries.GetSprint(r.ctx, id)
	return sprint, errors.MapDatabaseError(err)
}

func (r RequirementRepository) Create(requirement data.CreateRequirementParams) (data.Requirement, error) {
	created, err := r.queries.CreateRequirement(r.ctx, requirement)
	return created, errors.MapDatabaseError(err)
}

func (r RequirementRepository) Update(requirement data.UpdateRequirementParams) (data.Requirement, error) {
	updated, err := r.queries.UpdateRequirement(r.ctx, requirement)
	return updated, errors.MapDatabaseError(err)
}

func (r RequirementRepository) Delete(id uuid.UUID) error {
	return errors.MapDatabaseError(r.queries.DeleteRequirement(r.ctx, id))
}
