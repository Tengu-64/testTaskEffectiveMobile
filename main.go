package main

import (
	_ "api/docs"
	"api/handlers"
	"api/middleware"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

//@Title Тестовое API
//@description Реализация онлайн библиотеки песен

// @host localhost:5000
// @BasePath /songs

func main() {

	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	songs := router.Group("/songs")
	songs.Use(middleware.CheckConnectDb)
	{
		router.GET("/songs", handlers.GetSongs)
		router.POST("/songs/create", handlers.CreateSong)
		router.GET("songs/:id", handlers.GetSongById)
		router.PUT("/songs/change/:id", handlers.PutSong)
		router.DELETE("/songs/delete/:id", handlers.DeleteSong)
	}

	router.Run(":5000")
}
