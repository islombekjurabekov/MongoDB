package main

import (
	"MongoDB/controller"
	"MongoDB/services"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
)

var (
	server         *gin.Engine
	userService    services.UserService
	userController controller.UserController
	ctx            context.Context
	userCollection *mongo.Collection
	client         *mongo.Client
	err            error
)

func init() {
	ctx = context.TODO()

	conn := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err = mongo.Connect(ctx, conn)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")

	userCollection = client.Database("userdb").Collection("users")
	userService = services.NewUserService(userCollection, ctx)
	userController = controller.New(userService)
	server = gin.Default()

}

func main() {
	defer client.Disconnect(ctx)
	basePath := server.Group("v1")
	userController.RegisterUserRoutes(basePath)
	log.Fatal(server.Run(":3000"))
}
