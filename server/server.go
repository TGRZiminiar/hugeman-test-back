package server

import (
	"context"
	"log"
	"net/http"

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
	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// r.Run()

	// Cors
	// s.app.Use(cors.New(cors.Config{
	// 	AllowOrigins:  []string{"*"},
	// 	AllowMethods:  []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
	// 	AllowHeaders:  []string{"Origin"},
	// 	ExposeHeaders: []string{"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With"},
	// 	// AllowCredentials: true,
	// 	AllowOriginFunc: func(origin string) bool {
	// 		return origin == "http://localhost:3000"
	// 	},
	// 	MaxAge: 12 * time.Hour,
	// }))
	s.app.Use(cors.Default())

	// Body Limit 10 mb
	// s.app.Use(limits.RequestSizeLimiter(10))

	switch s.cfg.App.Name {
	case "todo":
		s.todoService()

	}

	s.httpListening()

}
