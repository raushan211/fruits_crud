package server

import "github.com/gin-gonic/gin"

func Server() {

	r := gin.Default()
	setupRotes(r)
	r.Run()
}
