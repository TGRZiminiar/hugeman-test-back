package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/TGRZiminiar/hugeman-test-back/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/mongo"
)

type (
	server struct {
		app *gin.Engine
		db  *mongo.Client
		cfg *config.Config
	}
)

func (s *server) httpListening() {
	if err := s.app.Run(s.cfg.App.Url); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Error: %v", err)
	}

}

func Start(pctx context.Context, cfg *config.Config, db *mongo.Client) {
	s := &server{
		app: gin.Default(),
		db:  db,
		cfg: cfg,
	}

	// Cors
	s.app.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
		AllowAllOrigins:  true,
	}))

	switch s.cfg.App.Name {
	case "todo":
		s.todoService()
	}

	s.httpListening()

}
