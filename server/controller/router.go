package controller

import (
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	r.GET("/", index)
	r.GET("/domains", domains)
	r.GET("/certificates", certificates)
	r.GET("/settings", settings)
}
