package handlers

import (
	"api/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Добавление песни
// @Description Добавление песни по получаемым параметрам в JSON
// @Accept json
// @Produce json
// @Param song body model.Song true "Информация о песне"
// @Router /songs/create [post]
func CreateSong(c *gin.Context) {
	db, _ := model.Connect()

	var song model.Song
	if err := c.BindJSON(&song); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	db.Create(&song)
	c.JSON(http.StatusOK, song)
}
