package handlers

import (
	"api/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Удаление песни
// @Description удвление песни по параметру запроса (id)
// @Produce      json
// @Param id path int true "ID песни"
// @Router /songs/delete/{id} [delete]
func DeleteSong(c *gin.Context) {
	id := c.Param("id")
	db, _ := model.Connect()

	var song model.Song
	if err := db.Delete(&song, id).Error; err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "песня успешно удалена"})

}
