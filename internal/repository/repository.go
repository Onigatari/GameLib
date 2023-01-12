package repository

import (
	"GameLib/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type GameLibOperations interface {
	GetAllList(ctx *gin.Context) ([]models.Game, error)
	CheckExist(game string, ctx *gin.Context) (bool, error)
	AddGameRequest(game string, ctx *gin.Context) error
	DeleteGameRequest(game string, ctx *gin.Context) error
	UpdateGameDoneRequest(game string, ctx *gin.Context) error
	GetRandomGames(ctx *gin.Context) (string, error)
	GetRandomListGames(ctx *gin.Context) ([]string, error)
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
	rows, err := r.db.Query("SELECT name, done FROM games ORDER BY done, LOWER(name)")
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

func (r *RequestPostgres) CheckExist(game string, ctx *gin.Context) (bool, error) {
	var request bool
	if err := r.db.QueryRow("SELECT EXISTS(SELECT name FROM games WHERE LOWER(name) = LOWER($1) LIMIT 1)", game).Scan(&request); err != nil {
		return false, err
	}

	if request {
		return false, nil
	}
	return true, nil
}

func (r *RequestPostgres) AddGameRequest(game string, ctx *gin.Context) error {
	if _, err := r.db.Exec("INSERT INTO games (name, done) VALUES ($1, false)", game); err != nil {
		return err
	}
	return nil
}

func (r *RequestPostgres) DeleteGameRequest(game string, ctx *gin.Context) error {
	if _, err := r.db.Exec("DELETE FROM games WHERE LOWER(name) = LOWER($1)", game); err != nil {
		return err
	}
	return nil
}

func (r *RequestPostgres) UpdateGameDoneRequest(game string, ctx *gin.Context) error {
	if _, err := r.db.Exec("UPDATE games SET DONE = NOT DONE WHERE name = $1", game); err != nil {
		return err
	}
	return nil
}

func (r *RequestPostgres) GetRandomGames(ctx *gin.Context) (string, error) {
	var result string
	if err := r.db.QueryRow("SELECT name FROM games WHERE done != TRUE ORDER BY random() LIMIT 1").Scan(&result); err != nil {
		return "", err
	}
	return result, nil
}

func (r *RequestPostgres) GetRandomListGames(ctx *gin.Context) ([]string, error) {
	rows, err := r.db.Query("SELECT name FROM games WHERE done != TRUE ORDER BY random() LIMIT 20")
	if err != nil {
		return []string{}, err
	}
	defer rows.Close()

	var gameList []string

	for rows.Next() {
		var tmp string
		if err := rows.Scan(&tmp); err != nil {
			return gameList, err
		}
		gameList = append(gameList, tmp)
	}

	if err = rows.Err(); err != nil {
		return gameList, err
	}
	return gameList, nil
}
