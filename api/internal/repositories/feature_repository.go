package repositories

import (
	"context"

	"github.com/bjvanbemmel/benkyou/internal/data"
	"github.com/bjvanbemmel/benkyou/internal/database"
	"github.com/bjvanbemmel/benkyou/internal/errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type FeatureRepository struct {
	conn    *pgxpool.Conn
	ctx     context.Context
	queries *data.Queries
}

func NewFeatureRepository(ctx context.Context) (FeatureRepository, error) {
	conn, err := database.New(context.Background())
	if err != nil {
		return FeatureRepository{}, err
	}

	return FeatureRepository{
		conn:    conn,
		ctx:     ctx,
		queries: data.New(conn),
	}, nil
}

func (f *FeatureRepository) CloseConn() error {
	return f.conn.Conn().Close(f.ctx)
}

func (f FeatureRepository) Index() ([]data.Sprint, error) {
	sprints, err := f.queries.ListSprints(f.ctx)
	return sprints, errors.MapDatabaseError(err)
}

func (f FeatureRepository) Get(id uuid.UUID) (data.Sprint, error) {
	sprint, err := f.queries.GetSprint(f.ctx, id)
	return sprint, errors.MapDatabaseError(err)
}

func (f FeatureRepository) Create(feature data.CreateFeatureParams) (data.Feature, error) {
	created, err := f.queries.CreateFeature(f.ctx, feature)
	return created, errors.MapDatabaseError(err)
}

func (f FeatureRepository) Update(feature data.UpdateFeatureParams) (data.Feature, error) {
	updated, err := f.queries.UpdateFeature(f.ctx, feature)
	return updated, errors.MapDatabaseError(err)
}

func (f FeatureRepository) Delete(id uuid.UUID) error {
	return errors.MapDatabaseError(f.queries.DeleteFeature(f.ctx, id))
}
