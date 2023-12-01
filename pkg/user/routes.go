package user

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shakezidin/middleware"
	"github.com/shakezidin/pkg/user/handler"
	userpb "github.com/shakezidin/pkg/user/pb"
	"github.com/shakezidin/pkg/config"
)

type User struct {
	cnfg   *config.Configure
	client userpb.UserServiceClient
}

func NewUserRoute(c *gin.Engine, cnfg *config.Configure) {
	Client, err := ClientDial(*cnfg)
	if err != nil {
		log.Fatalf("error Not connected with gRPC server, %v", err.Error())
	}

	userHandler := User{
		cnfg:   cnfg,
		client: Client,
	}

	apiUser := c.Group("/api/user")
	{
		//* Logging in
		apiUser.POST("/login", userHandler.Login)
		apiUser.POST("/create/user",userHandler.SignupUser)
	}
}

func (a *User) UserAuthenticate(ctx *gin.Context) {
	email, err := middleware.ValidateTocken(ctx, *a.cnfg, "user")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error":  err.Error(),
			"status": http.StatusUnauthorized,
		})
		return
	}
	ctx.Set("registered_email", email)
	ctx.Next()

}

func (a *User) Login(c *gin.Context) {
	handler.UserLoginHandler(c, a.client, "Admin")
}

func (a *User) SignupUser(c *gin.Context) {
	handler.CreateUserHandler(c, a.client)
}
