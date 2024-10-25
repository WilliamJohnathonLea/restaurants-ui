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
