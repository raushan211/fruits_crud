package handlers

import (
	"fruits_crud/server/pkg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteFruitHandler(c *gin.Context) {

	id, err := c.Params.Get("id")

	if err == false {
		res := gin.H{
			"error": "Something Went Wrong",
		}
		c.JSON(http.StatusBadGateway, res)
	}
	id_int, _ := strconv.Atoi(id)
	result := pkg.DeleteFruitsService(id_int)

	if result == false {
		res := gin.H{
			"error": "Something went wrong",
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := gin.H{
		"error": "Deleted Successfully",
	}
	c.JSON(http.StatusOK, res)
	return

}
