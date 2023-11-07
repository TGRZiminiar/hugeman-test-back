package main

import (
	"context"
	"log"
	"os"

	"github.com/TGRZiminiar/hugeman-test-back/config"
	"github.com/TGRZiminiar/hugeman-test-back/pkg/database"
	"github.com/TGRZiminiar/hugeman-test-back/server"
)

func main() {
	ctx := context.Background()
	_ = ctx

	cfg := config.LoadConfig(func() string {
		if len(os.Args) < 1 {
			log.Fatal("Error: .env path is invalid")
		}
		return os.Args[1]
	}())

	db := database.DbConn(ctx, &cfg)
	defer db.Disconnect(ctx)

	server.Start(ctx, &cfg, db)
}
