package todorepository

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/TGRZiminiar/hugeman-test-back/modules/todo"
	"github.com/TGRZiminiar/hugeman-test-back/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	TodoServiceRepository interface {
		FindOneTodo(pctx context.Context, todoId string) (*todo.Todo, error)
		InsertOneTodo(pctx context.Context, req *todo.Todo) (primitive.ObjectID, error)
		DeleteOneTodo(pctx context.Context, todoId string) (int64, error)
		FindManyTodo(pctx context.Context) ([]*todo.Todo, error)
		UpdateOneTodo(pctx context.Context, req *todo.TodoShowcase) error
	}

	todorepository struct {
		db *mongo.Client
	}
)

func NewTodoRepository(db *mongo.Client) TodoServiceRepository {
	return &todorepository{
		db: db,
	}
}

func (r *todorepository) todoDbConn(pctx context.Context) *mongo.Database {
	return r.db.Database("todo_db")
}

func (r *todorepository) FindOneTodo(pctx context.Context, todoId string) (*todo.Todo, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()
	db := r.todoDbConn(ctx)
	col := db.Collection("todos")

	result := new(todo.Todo)
	if err := col.FindOne(
		ctx,
		bson.M{"_id": utils.ConvertToObjectId(todoId)},
		options.FindOne().SetProjection(bson.M{
			"_id":         1,
			"title":       1,
			"description": 1,
			"created_at":  1,
			"updated_at":  1,
			"image":       1,
			"status":      1,
		}),
	).Decode(result); err != nil {
		log.Printf("Error: FindOneTodo: %s", err.Error())
		return nil, errors.New("error: find one todo not found")
	}
	return result, nil

}

func (r *todorepository) FindManyTodo(pctx context.Context) ([]*todo.Todo, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.todoDbConn(ctx)
	col := db.Collection("todos")

	// Option
	opts := make([]*options.FindOptions, 0)

	opts = append(opts, options.Find().SetSort(bson.D{{"_id", 1}}))
	// opts = append(opts, options.Find().SetLimit(int64(req.Limit)))

	cursors, err := col.Find(ctx, bson.M{}, opts...)
	if err != nil {
		log.Printf("Error: FindManyToDo failed: %s", err.Error())
		return nil, errors.New("error: todo list not found")
	}

	results := make([]*todo.Todo, 0)
	for cursors.Next(ctx) {
		result := new(todo.Todo)
		if err := cursors.Decode(result); err != nil {
			log.Printf("Error: FindManyToDo failed: %s", err.Error())
			return nil, errors.New("error: todo list not found")
		}

		results = append(results, result)
	}

	return results, nil

}

func (r *todorepository) UpdateOneTodo(pctx context.Context, req *todo.TodoShowcase) error {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()
	db := r.todoDbConn(ctx)
	col := db.Collection("todos")

	_, err := col.UpdateOne(
		pctx,
		bson.M{"_id": utils.ConvertToObjectId(req.Id)},
		bson.M{
			"$set": bson.M{
				"title":       req.Title,
				"description": req.Description,
				"image":       req.Image,
				"status":      req.Status,
				"updated_at":  req.UpdatedAt,
			}},
	)
	if err != nil {
		log.Printf("Error: Update One Player Credentail %s", err.Error())
		return errors.New("error: player credentail not found")
	}
	return nil
}

func (r *todorepository) InsertOneTodo(pctx context.Context, req *todo.Todo) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()
	db := r.todoDbConn(ctx)
	col := db.Collection("todos")

	todoId, err := col.InsertOne(ctx, req)
	if err != nil {
		log.Panicf("Error: InsertOneTodo Failed: %s", err.Error())
		return primitive.NilObjectID, errors.New("error: insert one todo failed")
	}

	return todoId.InsertedID.(primitive.ObjectID), nil
}

func (r *todorepository) DeleteOneTodo(pctx context.Context, todoId string) (int64, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.todoDbConn(ctx)
	col := db.Collection("todos")
	result, err := col.DeleteOne(ctx, bson.M{"_id": utils.ConvertToObjectId(todoId)})
	if err != nil {
		log.Printf("Error: DeleteOneTodo failed: %s", err.Error())
		return -1, errors.New("error: DeleteOneTodo failed")
	}

	log.Printf("Delete Result: %v", result)
	return result.DeletedCount, nil
}
