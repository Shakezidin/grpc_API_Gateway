package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/shakezidin/pkg/DTO"
	adminpb "github.com/shakezidin/pkg/admin/adminpb"
)

func AdminLoginHandler(c *gin.Context, client adminpb.AdminServiceClient, role string) {
	var admin DTO.AdminLogin
	if err := c.BindJSON(&admin); err != nil {
		log.Printf("error binding JSON")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}
	validate := validator.New()
	err := validate.Struct(admin)
	if err != nil {
		log.Printf("Validation error")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  "Validation error",
		})
	}
	ctx := context.Background()
	response, err := client.AdminLogin(ctx, &adminpb.LoginRequest{
		Username: admin.Username,
		Password: admin.Password,
		Role:     role,
	})
	if err != nil {
		log.Printf("error logging in user %v err: %v", admin.Username, err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusAccepted,
		"message": fmt.Sprintf("%v logged in succesfully", admin.Username),
		"data":    response,
	})
}

func CreateUserHandler(c *gin.Context, client adminpb.AdminServiceClient) {
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
	ctx := context.Background()
	responce, err := client.CreateUser(ctx, &adminpb.User{
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

type payload struct {
	Username string `json:"username"`
}

func SearchUserHandler(c *gin.Context, client adminpb.AdminServiceClient) {
	var name payload
	if err := c.BindJSON(&name); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  err.Error(),
		})
		return
	}

	ctx := context.Background()
	result, err := client.SearchUser(ctx, &adminpb.UserRequest{
		Username: name.Username,
	})
	if err != nil {
		log.Printf("error logging in user %v err: %v", name.Username, err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Status":  http.StatusOK,
		"message": fmt.Sprintf("%v fetched successfully", name.Username),
		"data":    result,
	})

}

func DeleteUserHandler(c *gin.Context, client adminpb.AdminServiceClient) {
	idstr := c.Query("Id")
	if idstr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  "id param missing",
		})
		return
	}

	id, err := strconv.Atoi(idstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  err.Error(),
		})
		return
	}
	ctx := context.Background()
	result, err := client.DeleteUser(ctx, &adminpb.DeleteUserRequest{
		Id: uint64(id),
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Status":  http.StatusOK,
		"message": fmt.Sprintf("%v deleted successfully", idstr),
		"data":    result,
	})
}

func EditUserHandler(c *gin.Context, client adminpb.AdminServiceClient) {
	idstr := c.Query("id")
	id, _ := strconv.Atoi(idstr)
	var user DTO.User
	if err := c.BindJSON(&user); err != nil {
		log.Printf("Error binding user")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  "Binding error",
		})
		return
	}
	ctx := context.Background()
	responce, err := client.EditUser(ctx, &adminpb.User{
		Id:       uint64(id),
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
		"Message": fmt.Sprintf("%v user updated successfully", user.Username),
		"data":    responce,
	})
}
