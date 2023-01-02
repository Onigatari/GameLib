package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h *Handler) mainPage(c *gin.Context) {
	base, err := h.services.GetAllList(c)
	if err != nil {
		log.Fatalf("request invalid: %s", err)
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"gameList":   base,
		"randomGame": "",
	})
}

func (h *Handler) addGameRequest(c *gin.Context) {
	game := c.Request.FormValue("game")
	check, err := h.services.CheckExist(game, c)
	if err != nil {
		log.Fatalf("request invalid: %s", err)
	}

	if check && len(game) > 0 {
		if err = h.services.AddGameRequest(game, c); err != nil {
			log.Fatalf("request invalid: %s", err)
		}
	}

	c.Redirect(http.StatusFound, "/")
}

func (h *Handler) deleteGameRequest(c *gin.Context) {
	game := c.Request.FormValue("buttonDelete")
	if err := h.services.DeleteGameRequest(game, c); err != nil {
		log.Fatalf("request invalid: %s", err)
	}
	c.Redirect(http.StatusFound, "/")
}

func (h *Handler) updateGameDoneRequest(c *gin.Context) {
	game := c.Request.FormValue("updateDoneGame")
	if err := h.services.UpdateGameDoneRequest(game, c); err != nil {
		log.Fatalf("request invalid: %s", err)
	}
	c.Redirect(http.StatusFound, "/")
}

func (h *Handler) getRandomGames(c *gin.Context) {
	base, errAll := h.services.GetAllList(c)
	if errAll != nil {
		log.Fatalf("request invalid: %s", errAll)
	}

	randomGame, errRand := h.services.GetRandomGames(c)
	if errRand != nil {
		log.Fatalf("request invalid: %s", errRand)
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"gameList":   base,
		"randomGame": randomGame,
	})
}
