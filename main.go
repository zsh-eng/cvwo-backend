package main

import (
	"net/http"

	//   "github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

type LOGIN struct {
	USERNAME string `json:"username" binding:"required"`
	PASSWORD string `json:"password" binding:"required"`
}

type FORUMPOST struct {
	TITLE string  `json:"title"`
	BODY  string  `json:"body"`
	OWNER string `json:"owner"`
}

func main() {
	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Serve frontend static files
	router.LoadHTMLGlob("resources/dist/*.html")
	router.Static("/assets", "./resources/dist/assets")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// Setup route group for the API
	api := router.Group("/api")
	api.GET("/posts", PostsHandler)
	api.POST("/posts", NewPostHandler)
	api.POST("/login", LoginHandler)

	// Start and run the server
	router.Run(":3000")
}

func PostsHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World!",
	})
}

func NewPostHandler(c *gin.Context) {
	var forumPost FORUMPOST

	// Error handling
	if err := c.BindJSON(&forumPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error!",
		})
		return
	}

	c.JSON(http.StatusCreated, forumPost)
}

func LoginHandler(c *gin.Context) {
	var login LOGIN

	// Error handling
	if err := c.BindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Success!",
	})
}
