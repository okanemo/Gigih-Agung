package api

import (
	db "backend/db/sqlc"

	"github.com/gin-gonic/gin"
)

// import "github.com/gin-gonic/gin"

// Server serves HTTP requests.
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// TODO: add routes to router
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)

	server.router = router
	return server
}

// Start runs the HTTP server on the input address to start listening on API requests
// it takes address as input and error as output
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

// Gives an error response
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
