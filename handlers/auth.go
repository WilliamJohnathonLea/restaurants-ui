package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/WilliamJohnathonLea/restaurants-ui/types"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func LoginHandler(ctx *gin.Context) {
	sess := sessions.Default(ctx)
	var req types.LoginRequest

	if err := ctx.Bind(&req); err != nil {
		ctx.Redirect(http.StatusSeeOther, "/")
		return
	}

	url := os.Getenv("USER_SERVICE_HOST") + "/login"
	jsonBytes, err := json.Marshal(req)
	if err != nil {
		ctx.Redirect(http.StatusSeeOther, "/")
		return
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonBytes))
	if err != nil {
		ctx.Redirect(http.StatusSeeOther, "/")
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		ctx.Redirect(http.StatusSeeOther, "/")
		return
	}

	var loginResp types.LoginResponse
	bs, err := io.ReadAll(resp.Body)
	if err != nil {
		ctx.Redirect(http.StatusSeeOther, "/")
		return
	}
	err = json.Unmarshal(bs, &loginResp)
	if err != nil {
		ctx.Redirect(http.StatusSeeOther, "/")
		return
	}
	sess.Set("user_token", loginResp.Token)
	sess.Save()
	ctx.Redirect(http.StatusSeeOther, "/")
}

func SignupHandler(ctx *gin.Context) {
	var req types.SignupRequest

	if err := ctx.Bind(&req); err != nil {
		ctx.Redirect(http.StatusSeeOther, "/")
		return
	}

	url := os.Getenv("USER_SERVICE_HOST") + "/users"
	jsonBytes, err := json.Marshal(req)
	if err != nil {
		ctx.Redirect(http.StatusSeeOther, "/")
		return
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonBytes))
	if err != nil {
		ctx.Redirect(http.StatusSeeOther, "/")
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		ctx.Redirect(http.StatusSeeOther, "/")
		return
	}

	ctx.Redirect(http.StatusSeeOther, "/")
}

func LogoutHander(ctx *gin.Context) {
	sess := sessions.Default(ctx)
	sess.Clear()
	sess.Save()

	ctx.HTML(http.StatusOK, "logged-out.html", gin.H{
		"title":     "Logged out",
		"signedIn":  false,
	})
}
