package repositories

import (
	"context"

	"github.com/bjvanbemmel/benkyou/internal/data"
	"github.com/bjvanbemmel/benkyou/internal/database"
	"github.com/bjvanbemmel/benkyou/internal/errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SprintRepository struct {
	conn    *pgxpool.Conn
	ctx     context.Context
	queries *data.Queries
}

func NewSprintRepository(ctx context.Context) (SprintRepository, error) {
	conn, err := database.New(context.Background())
	if err != nil {
		return SprintRepository{}, err
	}

	return SprintRepository{
		conn:    conn,
		ctx:     ctx,
		queries: data.New(conn),
	}, nil
}

func (s *SprintRepository) CloseConn() error {
	return s.conn.Conn().Close(s.ctx)
}

func (s SprintRepository) Index() ([]data.Sprint, error) {
	sprints, err := s.queries.ListSprints(s.ctx)
	return sprints, errors.MapDatabaseError(err)
}

func (s SprintRepository) Get(id uuid.UUID) (data.Sprint, error) {
	sprint, err := s.queries.GetSprint(s.ctx, id)
	return sprint, errors.MapDatabaseError(err)
}

func (s SprintRepository) Create(sprint data.CreateSprintParams) (data.Sprint, error) {
	created, err := s.queries.CreateSprint(s.ctx, sprint)
	return created, errors.MapDatabaseError(err)
}

func (s SprintRepository) Update(sprint data.UpdateSprintParams) (data.Sprint, error) {
	updated, err := s.queries.UpdateSprint(s.ctx, sprint)
	return updated, errors.MapDatabaseError(err)
}

func (s SprintRepository) Delete(id uuid.UUID) error {
	return errors.MapDatabaseError(s.queries.DeleteSprint(s.ctx, id))
}
