package main

import (
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

	// Home
	router.GET("/", handlers.Homepage)
	// Login
	router.GET("/login", handlers.LoginPage)
	router.POST("/login", handlers.LoginHandler)
	// Signup
	router.GET("/signup", handlers.SignupPage)
	router.POST("/signup", handlers.SignupHandler)
	// Logout
	router.GET("/logout", handlers.LogoutHander)

	router.Run(":8080")
}
