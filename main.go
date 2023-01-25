// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
// 	"runtime"

// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	ConfigRuntime()
// 	StartWorkers()
// 	StartGin()
// }

// // ConfigRuntime sets the number of operating system threads.
// func ConfigRuntime() {
// 	nuCPU := runtime.NumCPU()
// 	runtime.GOMAXPROCS(nuCPU)
// 	fmt.Printf("Running with %d CPUs\n", nuCPU)
// }

// // StartWorkers start starsWorker by goroutine.
// func StartWorkers() {
// 	go statsWorker()
// }

// // StartGin starts gin web server with setting router.
// func StartGin() {
// 	gin.SetMode(gin.ReleaseMode)

// 	router := gin.New()
// 	router.Use(rateLimit, gin.Recovery())
// 	router.LoadHTMLGlob("resources/*.templ.html")
// 	router.Static("/static", "resources/static")
// 	router.GET("/", index)
// 	router.GET("/room/:roomid", roomGET)
// 	router.POST("/room-post/:roomid", roomPOST)
// 	router.GET("/stream/:roomid", streamRoom)

// 	port := os.Getenv("PORT")
// 	if port == "" {
// 		port = "8080"
// 	}
// 	if err := router.Run(":" + port); err != nil {
//         log.Panicf("error: %s", err)
// 	}
// }

package main

import (
	"net/http"

	//   "github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Serve frontend static files
	//   router.Use(static.Serve("/", static.LocalFile("./resources/static/index.html", true)))
	router.LoadHTMLGlob("resources/dist/*.html")
	router.Static("/assets", "./resources/dist/assets")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// Setup route group for the API
	api := router.Group("/api")
	api.GET("/posts", PostsHandler)

	// Start and run the server
	router.Run(":3000")
}
func PostsHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World!",
	})
}
