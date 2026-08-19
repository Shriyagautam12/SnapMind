package handler

import "github.com/gin-gonic/gin"

type AuthHandler struct{}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func (h *AuthHandler) Register(c *gin.Context) {
	// TODO(phase-1): POST /api/v1/auth/register
}

func (h *AuthHandler) Login(c *gin.Context) {
	// TODO(phase-1): POST /api/v1/auth/login
}
