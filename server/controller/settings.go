package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func settings(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"views/settings.html",
		gin.H{
			"title": "TPM",
		},
	)
}
