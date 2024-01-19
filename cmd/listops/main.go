package main

import (
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/slash2314/listops/internal/templates"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

//go:generate templ generate

// Embed static assets
//
//go:embed assets
var f embed.FS

func main() {
	route := gin.New()
	route.StaticFS("/static/", http.FS(f))
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
	listenAddress := "127.0.0.1:3000"
	listenAddressString := os.Getenv("ADDRESS")
	if listenAddressString != "" {
		listenAddress = listenAddressString
	}
	handleSignals()
	err := http.ListenAndServe(fmt.Sprintf("%s", listenAddress), route)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func handleSignals() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	go func() {
		select {
		case <-sigs:
			os.Exit(0)
		}
	}()
}
