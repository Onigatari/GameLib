package service

import (
	"GameLib/internal/models"
	"GameLib/internal/repository"
	"github.com/gin-gonic/gin"
)

type GameLibOperations interface {
	GetAllList(ctx *gin.Context) ([]models.Game, error)
}

type GameLibService struct {
	GameLibOperations
}

func NewService(repos *repository.AppRepository) *GameLibService {
	return &GameLibService{
		GameLibOperations: NewGameLibOperationsService(repos.GameLibOperations),
	}
}
