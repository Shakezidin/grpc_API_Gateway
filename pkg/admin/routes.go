package admin

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shakezidin/middleware"
	pb "github.com/shakezidin/pkg/admin/adminpb"
	"github.com/shakezidin/pkg/admin/handler"
	"github.com/shakezidin/pkg/config"
)

type Admin struct {
	cnfg   *config.Configure
	client pb.AdminServiceClient
}

func NewAdminRoute(c *gin.Engine, cnfg *config.Configure) {
	Client, err := ClientDial(*cnfg)
	if err != nil {
		log.Fatalf("error Not connected with gRPC server, %v", err.Error())
	}

	adminHandler := Admin{
		cnfg:   cnfg,
		client: Client,
	}

	apiAdmin := c.Group("/api/admin")
	{
		//* Logging in
		apiAdmin.POST("/login", adminHandler.Login)
		apiAdmin.POST("/create/user", adminHandler.AdminAuthenticate, adminHandler.CreateUser)
		apiAdmin.POST("/search/user", adminHandler.AdminAuthenticate, adminHandler.SearchUser)
		apiAdmin.GET("/delete/user", adminHandler.AdminAuthenticate, adminHandler.DeleteUser)
		apiAdmin.PATCH("/edit/user", adminHandler.AdminAuthenticate, adminHandler.EditUser)
	}
}

func (a *Admin) AdminAuthenticate(ctx *gin.Context) {
	email, err := middleware.ValidateTocken(ctx, *a.cnfg, "admin")
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

func (a *Admin) Login(c *gin.Context) {
	handler.AdminLoginHandler(c, a.client, "Admin")
}

func (a *Admin) CreateUser(c *gin.Context) {
	handler.CreateUserHandler(c, a.client)
}

func (a *Admin) SearchUser(c *gin.Context) {
	handler.SearchUserHandler(c, a.client)
}

func (a *Admin) DeleteUser(c *gin.Context) {
	handler.DeleteUserHandler(c, a.client)
}

func (a *Admin) EditUser(c *gin.Context) {
	handler.EditUserHandler(c, a.client)
}
