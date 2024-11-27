package handlers

import (
	"api/model"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// @Summary Песня
// @Description Информация песни по id
// @Produce      json
// @Param id path int true "ID песни"
// @Router /songs/{id} [get]
func GetSongById(c *gin.Context) {
	id := c.Param("id")
	db, _ := model.Connect()
	var song model.Song

	if err := db.Find(&song, id).Error; err != nil {
		log.Println(err)
		c.JSON(http.StatusNotFound, gin.H{"message": "песня не найдена"})
		return
	}
	textToCouplets := strings.Split(string(song.Text), "\n")

	c.JSON(http.StatusOK, gin.H{
		"Id":          song.Id,
		"Group":       song.Group,
		"Name":        song.Name,
		"ReleaseDate": song.ReleaseDate,
		"Text":        textToCouplets,
		"Link":        song.Link,
	})
}

func stringToCouplets(txt string) []string {
	return strings.Split(txt, "\n\n")
}
