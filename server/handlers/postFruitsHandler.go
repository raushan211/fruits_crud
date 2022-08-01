package handlers

import (
	"fruits_crud/server/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostFruitsHandler(c *gin.Context) {
	reqBody := pkg.Fruit{}
	err := c.Bind(&reqBody)

	if err != nil {
		res := gin.H{
			"error": "Something went wrong",
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	result := pkg.FruitInsertService(reqBody)

	if result == false {

		res := gin.H{
			"error": "Something went Wrong",
		}
		c.JSON(http.StatusBadRequest, res)
	}

	res := gin.H{
		"message": "Fruit Inserted",
	}
	c.JSON(http.StatusCreated, res)

}
