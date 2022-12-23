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

func (s *GameLibOperationsService) CheckExist(game string, ctx *gin.Context) (bool, error) {
	return s.repo.CheckExist(game, ctx)
}

func (s *GameLibOperationsService) AddGameRequest(game string, ctx *gin.Context) error {
	return s.repo.AddGameRequest(game, ctx)
}

func (s *GameLibOperationsService) DeleteGameRequest(game string, ctx *gin.Context) error {
	return s.repo.DeleteGameRequest(game, ctx)
}

func (s *GameLibOperationsService) UpdateGameDoneRequest(game string, ctx *gin.Context) error {
	return s.repo.UpdateGameDoneRequest(game, ctx)
}
