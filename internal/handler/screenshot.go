package handler

import (
	"github.com/gin-gonic/gin"

	"screenshot-search/internal/service"
)

type ScreenshotHandler struct {
	screenshots *service.ScreenshotService
}

func NewScreenshotHandler(screenshots *service.ScreenshotService) *ScreenshotHandler {
	return &ScreenshotHandler{screenshots: screenshots}
}

func (h *ScreenshotHandler) Sync(c *gin.Context) {
	// TODO(phase-1): POST /api/v1/screenshots/sync
}

func (h *ScreenshotHandler) SyncStatus(c *gin.Context) {
	// TODO(phase-1): GET /api/v1/screenshots/sync/status
}
