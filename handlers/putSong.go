package handlers

import (
	"api/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Изменение песни
// @Description измение информации песни по id
// @Accept  json
// @Produce json
// @Param id path int true "ID песни"
// @Param song body model.Song true "Информация о песне"
// @Router /songs/change/{id} [put]
func PutSong(c *gin.Context) {
	db, _ := model.Connect()

	id := c.Param("id")
	var updateSong model.Song

	if err := c.BindJSON(&updateSong); err != nil {
		c.JSON(http.StatusBadRequest, err)
		log.Println(err)
		return
	}
	if err := db.Model(&model.Song{}).Where("id=?", id).Updates(updateSong).Error; err != nil {
		c.JSON(http.StatusNotFound, err)
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "данные песни успешно изменены"})
}
