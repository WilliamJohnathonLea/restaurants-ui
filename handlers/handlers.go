package handlers

import (
	"net/http"

	"github.com/WilliamJohnathonLea/restaurants-ui/types"
	"github.com/gin-gonic/gin"
)

func LoginHandler(ctx *gin.Context) {
	var req types.LoginRequest

	if err := ctx.Bind(&req); err != nil {
		ctx.Redirect(http.StatusBadRequest, "/")
		return
	}

	ctx.Redirect(http.StatusSeeOther, "/")
}

func SignupHandler(ctx *gin.Context) {
	var req types.SignupRequest

	if err := ctx.Bind(&req); err != nil {
		ctx.Redirect(http.StatusBadRequest, "/")
		return
	}

	ctx.Redirect(http.StatusSeeOther, "/")
}
