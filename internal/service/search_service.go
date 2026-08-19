package service

import (
	"context"

	"screenshot-search/internal/model"
	"screenshot-search/internal/repository"
)

type SearchService struct {
	repo *repository.ScreenshotRepository
	ai   AIProvider
}

func NewSearchService(repo *repository.ScreenshotRepository, ai AIProvider) *SearchService {
	return &SearchService{repo: repo, ai: ai}
}

func (s *SearchService) Search(ctx context.Context, userID, query string) ([]*model.Screenshot, error) {
	// TODO(phase-3): embed query text, hybrid search (vector + OCR/full-text), rank
	return nil, nil
}
