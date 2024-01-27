package controller

import (
	"MongoDB/models"
	"MongoDB/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	UserService services.UserService
}

func New(userService services.UserService) UserController {
	return UserController{
		UserService: userService,
	}
}

func (uc *UserController) CreateUser(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	err = uc.UserService.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusBadGateway, err)
	}
	c.JSON(http.StatusOK, gin.H{
		"Message": "Successfully created new user",
	})
}

func (uc *UserController) GetUser(c *gin.Context) {
	userName := c.Param("name")
	user, err := uc.UserService.GetUser(&userName)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusOK, user)
}

func (uc *UserController) GetAllUsers(c *gin.Context) {
	users, err := uc.UserService.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusOK, err)
	}
	c.JSON(http.StatusOK, users)
}

func (uc *UserController) UpdateUser(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	err = uc.UserService.UpdateUser(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusOK, gin.H{
		"Message": "Successfully updated",
	})
}

func (uc *UserController) DeleteUser(c *gin.Context) {
	userName := c.Param("name")
	err := uc.UserService.DeleteUser(&userName)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusOK, gin.H{
		"Message": "Successfully deleted",
	})
}

func (uc *UserController) RegisterUserRoutes(r *gin.RouterGroup) {
	userRoute := r.Group("/user")
	userRoute.POST("/create", uc.CreateUser)
	userRoute.GET("/get/:name", uc.GetUser)
	userRoute.GET("/getAllUsers", uc.GetAllUsers)
	userRoute.PATCH("/update", uc.UpdateUser)
	userRoute.DELETE("/delete/:name", uc.DeleteUser)
}
