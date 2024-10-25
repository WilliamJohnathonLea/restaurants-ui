package handlers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

func Homepage(ctx *gin.Context) {
	signedIn := sessions.Default(ctx).Get("user_token") != nil
	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"title":    "Food as a Service",
		"signedIn": signedIn,
	})
}

func LoginPage(ctx *gin.Context) {
	signedIn := sessions.Default(ctx).Get("user_token") != nil
	ctx.HTML(http.StatusOK, "login.html", gin.H{
		"title":     "Login",
		"signedIn":  signedIn,
		"csrfToken": csrf.GetToken(ctx),
	})
}

func SignupPage(ctx *gin.Context) {
	signedIn := sessions.Default(ctx).Get("user_token") != nil
	ctx.HTML(http.StatusOK, "signup.html", gin.H{
		"title":     "Signup",
		"signedIn":  signedIn,
		"csrfToken": csrf.GetToken(ctx),
	})
}

func RestaurantsPage(ctx *gin.Context) {
	signedIn := sessions.Default(ctx).Get("user_token") != nil
	ctx.HTML(http.StatusOK, "restaurants.html", gin.H{
		"title":        "Restaurants",
		"signedIn":     signedIn,
	})
}

func RestaurantPage(ctx *gin.Context) {
	signedIn := sessions.Default(ctx).Get("user_token") != nil
	restaurantID := ctx.Param("id")
	ctx.HTML(http.StatusOK, "restaurant.html", gin.H{
		"title":        "Restaurant",
		"restaurantId": restaurantID,
		"signedIn":     signedIn,
	})
}
