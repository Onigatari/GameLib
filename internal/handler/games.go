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
		"gameList": base,
	})
}
