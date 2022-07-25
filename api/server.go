package api

import (
	db "github.com/alexmolly/simple_bank/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

//create new http server
func NewServer(store db.Store) *Server {

	var server *Server = &Server{store: store}
	var router *gin.Engine = gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)

	server.router = router

	return server
}

//start new http server
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
