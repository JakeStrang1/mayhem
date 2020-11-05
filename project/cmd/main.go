package main

import (
	"example.com/myNetwork/db"
	"example.com/myNetwork/user"
	"github.com/gin-gonic/gin"
)

func main() {
	dbClient, err := db.NewClient("localhost:27017/mayhem")
	if err != nil {
		panic(err)
	}
	defer dbClient.Close()

	users := db.NewUserRepository(dbClient)
	userController := user.NewController(users)
	userHandler := user.NewHandler(userController)

	r := gin.Default()

	userRoutes := r.Group("/users")
	userRoutes.POST("/sign_up", userHandler.SignUp)

	r.Run()
}
