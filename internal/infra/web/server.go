package web

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Engine *gin.Engine
	DBConn *sql.DB
}

func NewServer(dbConn *sql.DB) *Server {
	return &Server{
		Engine: gin.Default(),
		DBConn: dbConn,
	}
}

func (s *Server) Run(port string) {
	SetupRoutes(s.Engine, s.DBConn)

	log.Printf("server running on port: %s", port)
	if err := s.Engine.Run(":" + port); err != nil {
		log.Fatalf("failed to start server %v", err)
	}
}
