package controller

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AuthController struct {
	Db *sql.DB
}

type RegisterRequest struct {
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password"`
	ImgUrl string `json:"img_url"`
}

func (a *AuthController) Register(ctx *gin.Context) {

	var req = RegisterRequest{}

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})
		return
	}

	val := validator.New()
	err = val.Struct(req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Payload" : req,
	})
}