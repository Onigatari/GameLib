package repository

import (
	"GameLib/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type GameLibOperations interface {
	GetAllList(ctx *gin.Context) ([]models.Game, error)
}

type AppRepository struct {
	GameLibOperations
}

func NewRepository(db *sqlx.DB) *AppRepository {
	return &AppRepository{
		GameLibOperations: NewPostgres(db),
	}
}

func (r *RequestPostgres) GetAllList(ctx *gin.Context) ([]models.Game, error) {
	rows, err := r.db.Query("SELECT name, done FROM games ORDER BY done, name")
	if err != nil {
		return []models.Game{}, err
	}
	defer rows.Close()

	var gameList []models.Game

	for rows.Next() {
		var tmp models.Game
		if err := rows.Scan(&tmp.Name, &tmp.Done); err != nil {
			return gameList, err
		}
		gameList = append(gameList, tmp)
	}

	if err = rows.Err(); err != nil {
		return gameList, err
	}
	return gameList, nil
}
