package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/mehdieidi/bank/db/sqlc"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(sotre *db.Store) *Server {
	server := &Server{store: sotre}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)

	server.router = router

	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
