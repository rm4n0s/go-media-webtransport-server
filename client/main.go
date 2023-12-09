package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Cross-Origin-Opener-Policy", "same-origin")
	c.Header("Cross-Origin-Embedder-Policy", "require-corp")

	c.Next()
}
func main() {
	router := gin.Default()
	//router.Use(cors.Default())
	// corsConf := &cors.Config{
	// 	AllowAllOrigins: true,
	// }
	// corsConf.AddAllowHeaders()
	router.Use(CORSMiddleware)
	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/encoder", func(c *gin.Context) {
		c.HTML(http.StatusOK, "encoder.html", gin.H{})
	})

	router.GET("/player", func(c *gin.Context) {
		c.HTML(http.StatusOK, "player.html", gin.H{})
	})

	log.Println("Run https://localhost:8080 and webtransport at https://localhost:4433/webtransport")
	router.RunTLS(":8080", "../certs/certificate.pem", "../certs/certificate.key")
}
