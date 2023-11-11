package migration

import (
	"context"
	"log"

	"github.com/TGRZiminiar/hugeman-test-back/config"
	"github.com/TGRZiminiar/hugeman-test-back/modules/todo"
	"github.com/TGRZiminiar/hugeman-test-back/pkg/database"
	"github.com/TGRZiminiar/hugeman-test-back/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func todoDbConn(pctx context.Context, cfg *config.Config) *mongo.Database {
	return database.DbConn(pctx, cfg).Database("todo_db")
}

func TodoMigrate(pctx context.Context, cfg *config.Config) {
	db := todoDbConn(pctx, cfg)
	defer db.Client().Disconnect(pctx)

	col := db.Collection("todos")

	indexs, _ := col.Indexes().CreateMany(pctx, []mongo.IndexModel{
		{
			Keys: bson.D{{"_id", 1}},
		},
		{
			Keys:    bson.D{{"title", "text"}},
			Options: options.Index().SetDefaultLanguage("english"),
		},
	})

	for _, index := range indexs {
		log.Printf("Indexs: %s", index)
	}

	documents := func() []any {
		roles := []*todo.Todo{
			{
				Title:       "Migrate Title1",
				Description: "Migrate Description1",
				Image:       "data:image/png;base64,image1",
				Status:      "IN_PROGRESS",
				CreatedAt:   utils.LocalTime(),
				UpdatedAt:   utils.LocalTime(),
			},
			{
				Title:       "Migrate Title2",
				Description: "Migrate Description2",
				Image:       "data:image/png;base64,image1",
				Status:      "IN_PROGRESS",
				CreatedAt:   utils.LocalTime(),
				UpdatedAt:   utils.LocalTime(),
			},
			{
				Title:       "Migrate Title3",
				Description: "Migrate Description3",
				Image:       "data:image/png;base64,image1",
				Status:      "COMPLETE",
				CreatedAt:   utils.LocalTime(),
				UpdatedAt:   utils.LocalTime(),
			},
		}

		docs := make([]any, 0)
		for _, r := range roles {
			docs = append(docs, r)
		}
		return docs
	}()

	results, err := col.InsertMany(pctx, documents, nil)
	if err != nil {
		panic(err)
	}
	log.Println("Migrate todo completed: ", results)

}
