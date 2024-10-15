package controller

import (
	"net/http"

	Domains "github.com/Fazendaaa/tpm/internal/domains"
	"github.com/gin-gonic/gin"
)

func index(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"views/index.html",
		gin.H{
			"title":   "TPM",
			"domains": Domains.SubdomainTest,
		},
	)
}
