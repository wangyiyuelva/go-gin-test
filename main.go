package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("templates/*")
	router.MaxMultipartMemory = 8 << 20 // 8 MiB

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	router.POST("/", func(c *gin.Context) {
		file, err := c.FormFile("video")
		// Get the file
		if err != nil {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"error": "fail to upload video.",
			})
			return
		}
		// Save the file
		err = c.SaveUploadedFile(file, "assets/uploads/"+file.Filename)

		// Render the page
		c.HTML(http.StatusOK, "index.html", gin.H{
			"video": "/assets/uploads/" + file.Filename,
		})
	})

	router.Run() // listen and serve on 0.0.0.0:8080
}
