package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/TGRZiminiar/hugeman-test-back/config"
	"github.com/gin-contrib/cors"
	limits "github.com/gin-contrib/size"
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
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))
	// Body Limit 10 mb
	s.app.Use(limits.RequestSizeLimiter(10))

	switch s.cfg.App.Name {
	case "todo":
		s.todoService()

	}

	s.httpListening()

}
