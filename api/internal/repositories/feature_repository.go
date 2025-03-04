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
	conn    *pgxpool.Pool
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

func (f FeatureRepository) Close() {
	f.conn.Close()
}

func (f FeatureRepository) Index() ([]data.Feature, error) {
	features, err := f.queries.ListFeatures(f.ctx)
	return features, errors.MapDatabaseError(err)
}

func (f FeatureRepository) Get(id uuid.UUID) (data.Feature, error) {
	feature, err := f.queries.GetFeature(f.ctx, id)
	return feature, errors.MapDatabaseError(err)
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

func (f FeatureRepository) IncrementPositionForAllFeaturesBetweenAscending(from int32, to int32) error {
	return f.queries.IncrementPositionForAllFeaturesBetweenAscending(f.ctx, data.IncrementPositionForAllFeaturesBetweenAscendingParams{
		Position:   from,
		Position_2: to,
	})
}

func (f FeatureRepository) IncrementPositionForAllFeaturesBetweenDescending(from int32, to int32) error {
	return f.queries.IncrementPositionForAllFeaturesBetweenDescending(f.ctx, data.IncrementPositionForAllFeaturesBetweenDescendingParams{
		Position:   to,
		Position_2: from,
	})
}
