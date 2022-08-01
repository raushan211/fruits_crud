package server

import (
	"fruits_crud/server/handlers"

	"github.com/gin-gonic/gin"
)

func setupRotes(r *gin.Engine) {
	r.GET("/fruits", handlers.GetFruitsHandler)
	r.POST("/fruits", handlers.PostFruitsHandler)
	r.PUT("/fruits", handlers.UpdateFruitHandler)
	r.DELETE("/fruits/:id", handlers.DeleteFruitHandler)
}
