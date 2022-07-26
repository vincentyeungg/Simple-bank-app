package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/vincentyeungg/Simple-bank-app/db/sqlc"
)

// Server serves HTTP requets for our banking service
type Server struct {
	store *db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and sets up routing
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// add routes to router
	router.POST("/accounts", server.createAccount)

	server.router = router
	return server
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}