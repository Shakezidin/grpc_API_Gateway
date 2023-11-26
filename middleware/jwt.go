package middleware

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/shakezidin/pkg/config"
)

type Claims struct {
	Email string
	Role  string
	jwt.StandardClaims
}

func ValidateTocken(c *gin.Context, cnfg config.Configure, role string) (string, error) {
	bearerToken := c.GetHeader("Authorization")
	if bearerToken == "" {
		log.Fatal("tocken missing")
		return "", errors.New("bearer token missing")
	}
	claims := Claims{}
	token := string([]byte(bearerToken)[7:])
	parseToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(cnfg.SECRETKEY), nil
	})
	if err != nil {
		log.Fatal("error while passing token")
		return "", errors.New("error passing error")
	}

	if !parseToken.Valid {
		log.Print("invalid token")
		return "", errors.New("token invalid")
	}

	expTime := claims.ExpiresAt
	if expTime < time.Now().Unix() {
		log.Print("token Expired")
		return "", errors.New("token expired")
	}

	fmt.Println(claims)
	userRole := claims.Role
	fmt.Println(role)
	if userRole != role {
		log.Println("unauthorized user")
		return "", errors.New("unauthorized user")
	}
	return claims.Email, nil
}
