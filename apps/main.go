package main

import (
	"nbfriends/apps/config"
	"nbfriends/apps/controller"
	"nbfriends/apps/pkg/token"
	"nbfriends/apps/response"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {

	// Set the timezone environment variable to "Asia/Jakarta"
	os.Setenv("TZ", "Asia/Jakarta")

	// Establish a connection to the database using the ConnectDB function from the config package
	// If an error occurs, panic with the error message
	db, err := config.ConnectDB()
	if err != nil {
		// Panic with error message if database connection fails
		panic(err)
	}

	// Create a new Gin router instance with default settings
	router := gin.Default() // gin default sudah ada log nya langsung

	// Initialize an instance of the AuthController struct, passing the database connection (db) to its Db field
	authController := controller.AuthController{
		// Database connection instance
		Db: db,
	}

	// Create a new API route group for version 1
	v1 := router.Group("v1")

	// Define a route for the "ping" endpoint with the Ping handler function
	router.GET("/ping", Ping)

	// Create a new route group for authentication-related endpoints
	auth := v1.Group("auth")
	{
		auth.POST("register", authController.Register)
		auth.POST("login", authController.Login)
		auth.GET("profile", CheckAuth(), authController.Profile)
	}

	// Start the Gin server on port 4444
	router.Run(":4444")

}

// CORS enables Cross-Origin Resource Sharing (CORS) for the Gin server.
//
// The Access-Control-Allow-Origin header is set to "*", allowing requests
// from any origin. The Access-Control-Request-Method and
// Access-Control-Allow-Headers headers are set to allow GET, POST, PUT, PATCH,
// DELETE, and OPTIONS requests and Content-Type and Authorization headers.
func CORS() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Request-Method", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		ctx.Header("Access-Control-Allow-Headers", "Authorization, Content-Type")
		ctx.Next()
	}
}

// Ping responds with a 200 status code and a JSON response containing the message "OK".
func Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "OK",
	})
}

func CheckAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get the Authorization header
		header := ctx.GetHeader("Authorization")

		// Check if the header starts with "Bearer "
		if !strings.HasPrefix(header, "Bearer ") {
			// Create a ResponseAPI object with the appropriate status code and message
			resp := response.ResponseAPI{
				StatusCode: http.StatusUnauthorized,
				Message:    "UNAUTHORIZED",
			}

			// Abort the request with the created ResponseAPI object
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, resp)
			return
		}

		// Extract the token by removing "Bearer " prefix and trimming spaces
		tokenString := strings.TrimSpace(header[7:])

		// Validate the token with the ValidateToken function from the token package
		// If validation fails, create a ResponseAPI object with the appropriate status code and message
		// and abort the request with the created ResponseAPI object
		payload, err := token.ValidateToken(tokenString)
		if err != nil {
			resp := response.ResponseAPI{
				StatusCode: http.StatusUnauthorized,
				Message:    "INVALID TOKEN",
				Payload:    err.Error(),
			}
			ctx.AbortWithStatusJSON(resp.StatusCode, resp)
			return
		}

		// Set the "authId" key in the context with the value of the AuthId field of the payload
		ctx.Set("authId", payload.AuthId)

		// Continue to the next middleware/handler in the chain
		ctx.Next()
	}
}

// router.GET("/ping", func(ctx *gin.Context) {
// 	ctx.JSON(http.StatusOK, map[string]interface{}{
// 		"message" : "OK",
// 	})
// })

// manual tanpa router
// http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	json.NewEncoder(w).Encode(map[string]interface{}{
// 		"message" : "OK",
// 	})
// })

// log.Println("server running on port :4444")
// http.ListenAndServe(":4444", nil)
