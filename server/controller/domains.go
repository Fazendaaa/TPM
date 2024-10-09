package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func domains(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"views/domains.html",
		gin.H{
			"title": "TPM",
		},
	)
}
