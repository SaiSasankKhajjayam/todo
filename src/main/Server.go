package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func CreateServer() {

	rg := TodoHandlers()
	srv = &http.Server{
		Addr:         port,
		Handler:      rg,
	}

	go func() {
		log.Println("Listening on port ", port)
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("listen: %s\n", err)
		}
	}()
}

//TodoHandlers  HTTP routes for TODO APP
func TodoHandlers() *gin.Engine {
	rg := gin.Default()

	rg.GET("/", func(c *gin.Context){
		c.JSON(http.StatusAccepted,"welcomee")
		return
	})
	rg.GET("/read",ReadHandler)
	rg.POST("/create", CreateHandler)
	rg.PUT("/update", UpdateHandler)
	rg.DELETE("/delete", DeleteHandler)

	return rg
}
