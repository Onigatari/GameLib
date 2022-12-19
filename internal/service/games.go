package service

import (
	"GameLib/internal/models"
	"GameLib/internal/repository"
	"github.com/gin-gonic/gin"
)

type GameLibOperationsService struct {
	repo repository.GameLibOperations
}

func NewGameLibOperationsService(repo repository.GameLibOperations) *GameLibOperationsService {
	return &GameLibOperationsService{repo: repo}
}
func (s *GameLibOperationsService) GetAllList(ctx *gin.Context) ([]models.Game, error) {
	return s.repo.GetAllList(ctx)
}
