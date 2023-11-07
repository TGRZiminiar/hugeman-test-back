package main

import (
	"context"
	"log"
	"os"

	"github.com/TGRZiminiar/hugeman-test-back/config"
	"github.com/TGRZiminiar/hugeman-test-back/pkg/database/migration"
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

	switch cfg.App.Name {
	case "todo":
		migration.TodoMigrate(ctx, &cfg)

	}

}
