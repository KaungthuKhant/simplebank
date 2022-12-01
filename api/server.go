package api

import (
	db "github.com/KaungthuKhant/simplebank/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests for our banking service
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// add router to router
	router.POST("/accounts", server.createAccount) // CreateAccount needs to be a method because we need to get acces to the object in order to save new account to the database

	// ser router object to server.router
	server.router = router
	return server
}

// Start function run the HTTP server on a specific address
func (server *Server) Start(address string) error {
	// router field is private so it cannot be access outside of this package
	return server.router.Run(address)
}

// return the error as map[string]interface{}
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
