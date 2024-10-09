package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func index(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"views/index.html",
		gin.H{
			"title": "TPM",
		},
	)
}
