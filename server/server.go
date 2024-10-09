package server

import (
	"log"

	controller "github.com/Fazendaaa/tpm/server/controller"

	"github.com/gin-gonic/gin"
)

func Server() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Static("/css", "./server/static/src/assets/css")
	r.Static("/images", "./server/static/src/assets/images")
	r.Static("/scss", "./server/static/src/assets/scss")
	r.Static("/js", "./server/static/src/assets/js")
	r.StaticFile("/favicon.ico", "./server/static/src/assets/images/favicon/favicon.ico")
	r.Static("/vendor", "./server/static/vendor")

	r.LoadHTMLGlob("./server/templates/**/*")
	controller.Router(r)

	log.Println("Server started")
	r.Run()
}
