package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/slash2314/listops/internal/templates"
	"log"
	"strconv"
)

type Op func(a, b []string) []string

func calcDiffRoute(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		fmt.Println(err.Error())
		c.Status(500)
	}
	delimiter := "\n"
	switch c.Request.PostFormValue("delimiter") {
	case "comma":
		delimiter = ","

	}

	var op Op
	switch c.Request.PostFormValue("difftype") {

	case "adiffb":
		op = aDiffB
	case "bdiffa":
		op = bDiffA
	case "inter":
		op = intersection
	case "union":
		op = union
	default:
		fmt.Println("Unknown diff type", c.Request.PostFormValue("difftype"))
		c.Status(500)
		return
	}
	err = setOperations(c, delimiter, op)
	if err != nil {
		log.Println("Error rendering Set Operations AJAX responder")
		c.Status(500)
	}
}

func calcChunksRoute(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		fmt.Println(err.Error())
		c.Status(500)
	}
	serialList := c.Request.PostFormValue("seriallist")
	rawSize := c.Request.PostFormValue("size")
	var chunkSize int
	if rawSize == "" {
		chunkSize = 20
	} else {
		chunkSize, err = strconv.Atoi(rawSize)
		if err != nil {
			chunkSize = 20
		}
	}
	outputDelimiter := ", "
	switch c.Request.PostFormValue("output-delimiter") {
	case "without":
		outputDelimiter = ","
	}

	delimiter := "\n"
	switch c.Request.PostFormValue("delimiter") {
	case "comma":
		delimiter = ","

	}
	chunkedInput := Chunk(serialList, delimiter, chunkSize)
	component := templates.ChunkResponder(chunkedInput, outputDelimiter)
	err = component.Render(c.Request.Context(), c.Writer)
	if err != nil {
		log.Println("Error rendering Chunks AJAX responder")
		c.Status(500)
	}
}

func apiRoutes(route *gin.Engine) {
	route.POST("/calcchunks", func(c *gin.Context) {
		calcChunksRoute(c)
	})
	route.POST("/calcdiff", func(c *gin.Context) {
		calcDiffRoute(c)

	})
}
