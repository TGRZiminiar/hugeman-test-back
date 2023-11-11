package todousecase

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/TGRZiminiar/hugeman-test-back/modules/todo"
	todorepository "github.com/TGRZiminiar/hugeman-test-back/modules/todo/todoRepository"
)

type (
	TodoUseCaseService interface {
		InsertOneTodo(pctx context.Context, req *todo.CreateTodoReq) (*todo.TodoShowcase, error)
		FindOneTodo(pctx context.Context, todoId string) (*todo.TodoShowcase, error)
		DeleteOneTodo(pctx context.Context, todoId string) (int64, error)
		FindManyTodo(pctx context.Context, page, limit int, sort string) ([]*todo.TodoShowcase, error)
		UpdateOneTodo(pctx context.Context, req *todo.UpdateTodoReq) (*todo.TodoShowcase, error)
		SearchTodo(pctx context.Context, text string) ([]*todo.TodoShowcase, error)
	}

	todoUsecase struct {
		todoRepository todorepository.TodoServiceRepository
	}
)

func NewTodoUsecase(todoRepository todorepository.TodoServiceRepository) TodoUseCaseService {
	return &todoUsecase{todoRepository: todoRepository}
}

func (u *todoUsecase) InsertOneTodo(pctx context.Context, req *todo.CreateTodoReq) (*todo.TodoShowcase, error) {
	todoId, err := u.todoRepository.InsertOneTodo(pctx, &todo.Todo{
		Title:       req.Title,
		Description: req.Description,
		Image:       req.Image,
		Status:      req.Status,
	})
	if err != nil {
		return nil, err
	}

	return u.FindOneTodo(pctx, todoId.Hex())
}

func (u *todoUsecase) UpdateOneTodo(pctx context.Context, req *todo.UpdateTodoReq) (*todo.TodoShowcase, error) {
	err := u.todoRepository.UpdateOneTodo(pctx, &todo.TodoShowcase{
		Id:          req.Id,
		Title:       req.Title,
		Description: req.Description,
		Image:       req.Image,
		Status:      req.Status,
	})
	if err != nil {
		return nil, err
	}

	return u.FindOneTodo(pctx, req.Id)
}

func (u *todoUsecase) FindOneTodo(pctx context.Context, todoId string) (*todo.TodoShowcase, error) {
	result, err := u.todoRepository.FindOneTodo(pctx, todoId)
	if err != nil {
		return nil, err
	}

	loc, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		log.Printf("Error: FindOneTodo Load Location Failed: %s", err.Error())
		return nil, errors.New("error: failed to load location")
	}

	return &todo.TodoShowcase{
		Id:          result.Id.Hex(),
		Title:       result.Title,
		Description: result.Description,
		CreatedAt:   result.CreatedAt.In(loc),
		UpdatedAt:   result.UpdatedAt.In(loc),
		Image:       "data:image/png;base64," + result.Image,
		Status:      result.Status,
	}, nil
}

func (u *todoUsecase) FindManyTodo(pctx context.Context, page, limit int, sort string) ([]*todo.TodoShowcase, error) {

	results, err := u.todoRepository.FindManyTodo(pctx, page, limit, sort)
	if err != nil {
		return make([]*todo.TodoShowcase, 0), err
	}

	if len(results) == 0 {
		return make([]*todo.TodoShowcase, 0), errors.New("error: no todo list found")
	}

	todos := make([]*todo.TodoShowcase, 0)

	loc, _ := time.LoadLocation("Asia/Bangkok")

	for _, v := range results {
		todos = append(todos, &todo.TodoShowcase{
			Id:          v.Id.Hex(),
			Title:       v.Title,
			Description: v.Description,
			Image:       "data:image/png;base64," + v.Image,
			Status:      v.Status,
			CreatedAt:   v.CreatedAt.In(loc),
			UpdatedAt:   v.UpdatedAt.In(loc),
		})
	}

	return todos, nil

}

func (u *todoUsecase) DeleteOneTodo(pctx context.Context, todoId string) (int64, error) {
	count, err := u.todoRepository.DeleteOneTodo(pctx, todoId)
	if err != nil {
		return -1, err
	}
	return count, nil
}

func (u *todoUsecase) SearchTodo(pctx context.Context, text string) ([]*todo.TodoShowcase, error) {

	results, err := u.todoRepository.SearchTodo(pctx, text)
	if err != nil {
		return make([]*todo.TodoShowcase, 0), err
	}
	if err != nil {
		return make([]*todo.TodoShowcase, 0), err
	}

	if len(results) == 0 {
		return make([]*todo.TodoShowcase, 0), errors.New("error: no todo list found")
	}

	todos := make([]*todo.TodoShowcase, 0)

	loc, _ := time.LoadLocation("Asia/Bangkok")

	for _, v := range results {
		todos = append(todos, &todo.TodoShowcase{
			Id:          v.Id.Hex(),
			Title:       v.Title,
			Description: v.Description,
			Image:       "data:image/png;base64," + v.Image,
			Status:      v.Status,
			CreatedAt:   v.CreatedAt.In(loc),
			UpdatedAt:   v.UpdatedAt.In(loc),
		})
	}

	return todos, nil
}
