package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Hello",
		})
	})

	// Login
	router.GET("/login", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login.html", gin.H{
			"title": "Login",
			"check": uuid.NewString(),
		})
	})

	router.POST("/login", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusSeeOther, "/")
	})

	// Signup
	router.GET("/signup", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "signup.html", gin.H{
			"title": "Signup",
			"check": uuid.NewString(),
		})
	})

	router.POST("/signup", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusSeeOther, "/")
	})

	router.Run(":8080")
}
