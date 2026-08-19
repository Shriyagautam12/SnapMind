package handler

import (
	"github.com/gin-gonic/gin"

	"screenshot-search/internal/service"
)

type SearchHandler struct {
	search *service.SearchService
}

func NewSearchHandler(search *service.SearchService) *SearchHandler {
	return &SearchHandler{search: search}
}

func (h *SearchHandler) Search(c *gin.Context) {
	// TODO(phase-3): POST /api/v1/search
}
