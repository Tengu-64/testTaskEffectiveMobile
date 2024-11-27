package handlers

import (
	"api/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Все песни
// @Description Просмотр всех песен с пагинацией и фильтрацией данных
// @Accept  json
// @Produce json
// @Param page query int false "Номер страницы" default(1)
// @Param limit query int false "Количество элементов на странице" default(10)
// @Param group query string false "Название группы"
// @Param releasedate query string false "Дата релиза"
// @Router /songs [get]
func GetSongs(c *gin.Context) {
	db, _ := model.Connect()

	var songs []model.Song
	var total int64

	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")
	group := c.Query("group")
	releaseDate := c.Query("releasedate")

	pageInt, _ := strconv.Atoi(page)
	limitInt, _ := strconv.Atoi(limit)

	offset := (pageInt - 1) * limitInt
	query := db.Limit(limitInt).Offset(offset)

	if group != "" {
		query = query.Where("\"group\" ILIKE ?", "%"+group+"%")
	}
	if releaseDate != "" {
		query = query.Where("\"release_date\" ILIKE ?", "%"+releaseDate+"%")
	}

	query.Find(&songs).Count(&total)

	c.JSON(http.StatusOK, gin.H{
		"data":  songs,
		"total": total,
		"page":  pageInt,
		"limit": limitInt,
	})

}
