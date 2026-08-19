package worker

import (
	"context"

	"screenshot-search/internal/repository"
	"screenshot-search/internal/service"
	"screenshot-search/internal/storage"
)

// IndexingWorker pulls pending screenshots, runs AI extraction, and deletes
// the source image from object storage on success. Jobs must be idempotent:
// retries must never create duplicate rows.
type IndexingWorker struct {
	repo    *repository.ScreenshotRepository
	storage storage.ObjectStorage
	ai      service.AIProvider
}

func NewIndexingWorker(repo *repository.ScreenshotRepository, storage storage.ObjectStorage, ai service.AIProvider) *IndexingWorker {
	return &IndexingWorker{repo: repo, storage: storage, ai: ai}
}

func (w *IndexingWorker) Run(ctx context.Context) error {
	// TODO(phase-2): job queue + worker pool, bounded AI/upload concurrency
	return nil
}
