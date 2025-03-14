package main

import (
	"nbfriends/apps/config"
	"nbfriends/apps/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	db ,err := config.ConnectDB()
	if err != nil {
		panic(err)
	}

	router := gin.Default() // gin default sudah ada log nya langsung

	authController := controller.AuthController{
		Db: db,
	}

	// router.GET("/ping", func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusOK, map[string]interface{}{
	// 		"message" : "OK",
	// 	})
	// })

	router.GET("/ping", Ping)
	router.POST("/v1/auth/register", authController.Register)

	router.Run(":4444")

	// manual tanpa router
	// http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Set("Content-Type", "application/json")

	// 	json.NewEncoder(w).Encode(map[string]interface{}{
	// 		"message" : "OK",
	// 	})
	// })

	// log.Println("server running on port :4444")
	// http.ListenAndServe(":4444", nil)
}

func Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"message" : "OK",
	})
}