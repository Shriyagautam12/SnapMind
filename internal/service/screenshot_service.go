package service

import (
	"context"

	"screenshot-search/internal/repository"
	"screenshot-search/internal/storage"
)

type ScreenshotService struct {
	repo    *repository.ScreenshotRepository
	storage storage.ObjectStorage
}

func NewScreenshotService(repo *repository.ScreenshotRepository, storage storage.ObjectStorage) *ScreenshotService {
	return &ScreenshotService{repo: repo, storage: storage}
}

// Sync registers screenshot references for a user and returns upload URLs
// for any not already indexed (matched by content hash).
func (s *ScreenshotService) Sync(ctx context.Context, userID string, deviceAssetIDs, contentHashes []string) (map[string]string, error) {
	// TODO(phase-1): dedup against content_hash, create pending rows, return upload URLs
	return nil, nil
}

func (s *ScreenshotService) SyncStatus(ctx context.Context, userID string) (map[string]int, error) {
	// TODO(phase-1): wire to repo.CountByStatus
	return nil, nil
}
