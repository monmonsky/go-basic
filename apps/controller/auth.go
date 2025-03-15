package controller

import (
	"database/sql"
	"nbfriends/apps/pkg/token"
	"nbfriends/apps/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	Db *sql.DB
}

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password"`
	ImgUrl   string `json:"img_url"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password"`
}

type Auth struct {
	Id       int
	Email    string
	Password string
}

var (
	queryCreate = `
		INSERT INTO auth (email, password, img_url)
		VALUES ($1, $2, $3)
	`

	queryFindByEmail = `
		SELECT id, email, password
		FROM auth
		WHERE email = $1
	`
)

func (a *AuthController) Register(ctx *gin.Context) {

	// bind data from request body to struct
	var req = RegisterRequest{}

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		// return error if binding failed
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// validate request with validator
	val := validator.New()
	err = val.Struct(req)
	if err != nil {
		// return error if validation failed
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// hash password with bcrypt
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		// return error if hashing failed
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// replace password with hashed password
	req.Password = string(hash)

	// prepare query to create new user
	stmt, err := a.Db.Prepare(queryCreate)
	if err != nil {
		// return error if query preparation failed
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// execute query to create new user
	_, err = stmt.Exec(
		req.Email,
		req.Password,
		req.ImgUrl,
	)
	if err != nil {
		// return error if query execution failed
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// prepare response
	resp := response.ResponseAPI{
		StatusCode: http.StatusCreated,
		Message:    "CREATED SUCCESS",
	}

	// return response with http.StatusCreated
	ctx.JSON(resp.StatusCode, resp)
}

func (a *AuthController) Login(ctx *gin.Context) {
	// bind data from request body to struct
	var req = LoginRequest{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		// return error if binding failed
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// prepare query to find user by email
	stmt, err := a.Db.Prepare(queryFindByEmail)
	if err != nil {
		// return error if prepare failed
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// execute query and get result
	row := stmt.QueryRow(req.Email)

	// create new struct to store result
	var auth = Auth{}

	// scan result to struct
	err = row.Scan(&auth.Id, &auth.Email, &auth.Password)
	if err != nil {
		// return error if scan failed
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// compare password from request with password from db
	err = bcrypt.CompareHashAndPassword([]byte(auth.Password), []byte(req.Password))
	if err != nil {
		// return error if password not match
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	// create payload token
	tok := token.PayloadToken{
		AuthId: auth.Id,
	}

	// generate token
	tokenString, err := token.GenerateToken(&tok)
	if err != nil {
		// return error if generate token failed
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// create response
	resp := response.ResponseAPI{
		StatusCode: http.StatusOK,
		Message:    "LOGIN SUCCESS",
		Payload: gin.H{
			"token": tokenString,
		},
	}

	// return response
	ctx.JSON(resp.StatusCode, resp)
}

func (a *AuthController) Profile(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"id": ctx.GetInt("authId"),
	})
}
