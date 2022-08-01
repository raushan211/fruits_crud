package handlers

import (
	"fruits_crud/server/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetFruitsHandler(c *gin.Context) {
	data := pkg.GetFruitsService()

	res := gin.H{

		"data": data,
	}
	c.JSON(http.StatusOK, res)
	return

}
