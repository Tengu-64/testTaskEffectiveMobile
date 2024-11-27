package middleware

import (
	"api/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckConnectDb(c *gin.Context) {
	_, err := model.Connect()
	if err != nil {
		//log.Fatal(err)
		c.JSON(http.StatusInternalServerError, err)
		log.Println(err)
		c.Abort()
	}
}
