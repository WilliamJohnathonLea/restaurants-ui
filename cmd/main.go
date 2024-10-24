package main

import (
	"net/http"
	"os"

	"github.com/WilliamJohnathonLea/restaurants-ui/handlers"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	csrf "github.com/utrack/gin-csrf"
)

func main() {
	godotenv.Load()

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	store := cookie.NewStore([]byte(os.Getenv("SESSION_STORE_SECRET")))
	router.Use(sessions.Sessions("faas-session", store))

	router.Use(csrf.Middleware(csrf.Options{
		Secret: os.Getenv("CSRF_SECRET"),
	}))

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Hello",
		})
	})

	// Login
	router.GET("/login", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login.html", gin.H{
			"title": "Login",
			"csrfToken": csrf.GetToken(ctx),
		})
	})

	router.POST("/login", handlers.LoginHandler)

	// Signup
	router.GET("/signup", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "signup.html", gin.H{
			"title": "Signup",
			"csrfToken": csrf.GetToken(ctx),
		})
	})

	router.POST("/signup", handlers.SignupHandler)

	router.Run(":8080")
}
