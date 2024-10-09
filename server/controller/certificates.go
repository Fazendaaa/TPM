package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func certificates(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"views/certificates.html",
		gin.H{
			"title": "TPM",
		},
	)
}
