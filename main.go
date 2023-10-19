package main

import (
	"log"
	"net/http"

	"github.com/chtiwa/go_file_cloudinary/controllers"
	"github.com/chtiwa/go_file_cloudinary/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	r := gin.Default()
	r.Static("/assets", "./assets")
	r.LoadHTMLGlob("templates/*")
	r.MaxMultipartMemory = 8 << 20 // 8 Mib

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main page",
		})
	})

	r.POST("/", func(c *gin.Context) {
		// single file
		file, err := c.FormFile("file")
		log.Println(file.Filename)

		if err != nil {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"error": "Failed to upload image",
			})
		}

		// Upload the file to specific dst.
		err = c.SaveUploadedFile(file, "assets/uploads/"+file.Filename)

		if err != nil {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"error": "Failed to upload image",
			})
		}

		c.HTML(http.StatusOK, "index.html", gin.H{
			"image": "/assets/uploads/" + file.Filename,
		})
	})

	r.POST("/cloudinary/file-upload", controllers.FileUpload())
	// r.POST("/cloudinary/remote-upload", controllers.RemoteUpload())

	r.Run()
}
