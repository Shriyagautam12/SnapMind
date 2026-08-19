package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"screenshot-search/internal/model"
)

type ScreenshotRepository struct {
	db *pgxpool.Pool
}

func NewScreenshotRepository(db *pgxpool.Pool) *ScreenshotRepository {
	return &ScreenshotRepository{db: db}
}

func (r *ScreenshotRepository) Create(ctx context.Context, s *model.Screenshot) error {
	// TODO(phase-1): insert unprocessed screenshot row
	return nil
}

func (r *ScreenshotRepository) FindByContentHash(ctx context.Context, userID, contentHash string) (*model.Screenshot, error) {
	// TODO(phase-1): used for incremental indexing / dedup
	return nil, nil
}

func (r *ScreenshotRepository) UpdateIndexingResult(ctx context.Context, s *model.Screenshot) error {
	// TODO(phase-2): persist metadata + embedding, mark indexed
	return nil
}

func (r *ScreenshotRepository) CountByStatus(ctx context.Context, userID string) (map[model.IndexingStatus]int, error) {
	// TODO(phase-1): power GET /screenshots/sync/status
	return nil, nil
}

func (r *ScreenshotRepository) SearchByEmbedding(ctx context.Context, userID string, queryEmbedding []float32, limit int) ([]*model.Screenshot, error) {
	// TODO(phase-3): pgvector similarity search
	return nil, nil
}
