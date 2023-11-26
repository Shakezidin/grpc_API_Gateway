package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/shakezidin/pkg/DTO"
	pb "github.com/shakezidin/pkg/user/pb"
)

func UserLoginHandler(c *gin.Context, client pb.UserServiceClient, role string) {
	var user DTO.User
	if err := c.BindJSON(&user); err != nil {
		log.Printf("error binding JSON")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}
	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		log.Printf("Validation error")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  "Validation error",
		})
	}
	var ctx context.Context
	response, err := client.UserLogin(ctx, &pb.LoginRequest{
		Username: user.Username,
		Password: user.Password,
		Role:     role,
	})
	if err != nil {
		log.Printf("error logging in user %v err: %v", user.Username, err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusAccepted,
		"message": fmt.Sprintf("%v logged in succesfully", user.Username),
		"data":    response,
	})
}

func CreateUserHandler(c *gin.Context, client pb.UserServiceClient) {
	var user DTO.User
	if err := c.BindJSON(&user); err != nil {
		log.Printf("Error binding user")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  "Binding error",
		})
		return
	}
	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		log.Print("Validation error")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  "Validation error",
		})
		return
	}
	var ctx context.Context
	responce, err := client.UserSignup(ctx, &pb.SignupRequest{
		Username: user.Username,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	})
	if err != nil {
		log.Printf("error logging in user %v err: %v", user.Username, err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"Status":  http.StatusAccepted,
		"Message": fmt.Sprintf("%v created successfully", user.Username),
		"data":    responce,
	})
}
