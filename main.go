package main

import (
	"bwastartup/handler"
	"bwastartup/user"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	/**
	* Connecting to database
	 */
	dsn := "root:passw0rd@tcp(127.0.0.1:3306)/crowdfund?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	// repository instance
	userRepository := user.NewRepository(db)

	// service instance
	userService := user.NewService(userRepository)

	// handler instance
	userHandler := handler.NewUserHandler(userService)

	// init gin router
	router := gin.Default()

	// set api group for `/api/v1`
	api := router.Group("/api/v1")

	// list of router
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)

	// listen server on port 3001
	router.Run(":3001")

}
