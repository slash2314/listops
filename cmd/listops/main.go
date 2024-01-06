package main

import (
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/slash2314/listops/internal/templates"
	"log"
	"net/http"
	"os"
	"strconv"
)

//go:generate templ generate

// Embed static assets
//
//go:embed assets/*
var f embed.FS

func main() {
	route := gin.New()
	// route.Static("/assets/", "./assets/")
	route.StaticFS("/assets", http.FS(f))
	route.GET("/sets", func(c *gin.Context) {
		component := templates.SetsForm()
		err := component.Render(c.Request.Context(), c.Writer)
		if err != nil {
			log.Println("Error rendering sets template")
			c.Status(500)
		}

	})
	route.GET("/", func(c *gin.Context) {
		component := templates.MainPage()
		err := component.Render(c.Request.Context(), c.Writer)
		if err != nil {
			log.Println("Error rendering MainPage template")
			c.Status(500)
		}

	})
	route.GET("/chunks", func(c *gin.Context) {
		component := templates.ChunksForm()
		err := component.Render(c.Request.Context(), c.Writer)
		if err != nil {
			log.Println("Error rendering Chunks template")
			c.Status(500)
		}

	})
	apiRoutes(route)
	listenPort := 3000
	portString := os.Getenv("PORT")
	if portString != "" {
		if parsedPort, err := strconv.Atoi(portString); err == nil {
			listenPort = parsedPort
		}
	}
	err := http.ListenAndServe(fmt.Sprintf("127.0.0.1:%d", listenPort), route)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
