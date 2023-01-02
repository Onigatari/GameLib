package service

import (
	"GameLib/internal/models"
	"GameLib/internal/repository"
	"github.com/gin-gonic/gin"
)

type GameLibOperations interface {
	GetAllList(ctx *gin.Context) ([]models.Game, error)
	CheckExist(game string, ctx *gin.Context) (bool, error)
	AddGameRequest(game string, ctx *gin.Context) error
	DeleteGameRequest(game string, ctx *gin.Context) error
	UpdateGameDoneRequest(game string, ctx *gin.Context) error
	GetRandomGames(ctx *gin.Context) (string, error)
}

type GameLibService struct {
	GameLibOperations
}

func NewService(repos *repository.AppRepository) *GameLibService {
	return &GameLibService{
		GameLibOperations: NewGameLibOperationsService(repos.GameLibOperations),
	}
}
