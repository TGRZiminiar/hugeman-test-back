package migration

import (
	"context"
	"log"

	"github.com/TGRZiminiar/hugeman-test-back/config"
	"github.com/TGRZiminiar/hugeman-test-back/pkg/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func paymentDbConn(pctx context.Context, cfg *config.Config) *mongo.Database {
	return database.DbConn(pctx, cfg).Database("todo_db")
}

func TodoMigrate(pctx context.Context, cfg *config.Config) {
	db := paymentDbConn(pctx, cfg)
	defer db.Client().Disconnect(pctx)

	col := db.Collection("todo")

	// Insert Here
	results, err := col.InsertOne(pctx, bson.M{"offset": -1}, nil)
	if err != nil {
		panic(err)
	}
	log.Println("Migrate todo completed: ", results)
}
