package server

import "github.com/gin-gonic/gin"

type ServerStruct struct {
	R *gin.Engine
}

func (s *ServerStruct) StartServer(port string) {
	s.R.Run(":"+port)
}

func Server() *ServerStruct {
	engine := gin.Default()

	return &ServerStruct{
		R: engine,
	}
}
