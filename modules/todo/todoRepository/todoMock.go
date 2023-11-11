package todorepository

import (
	"context"

	"github.com/TGRZiminiar/hugeman-test-back/modules/todo"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TodoRepoMock struct {
	mock.Mock
}

func (m *TodoRepoMock) FindOneTodo(pctx context.Context, todoId string) (*todo.Todo, error) {
	args := m.Called(pctx, todoId)
	return args.Get(0).(*todo.Todo), args.Error(1)
}
func (m *TodoRepoMock) FindManyTodo(pctx context.Context, page, limit int, sort string) ([]*todo.Todo, error) {
	args := m.Called(pctx, page, limit, sort)
	return args.Get(0).([]*todo.Todo), args.Error(1)
}

func (m *TodoRepoMock) UpdateOneTodo(pctx context.Context, req *todo.TodoShowcase) error {
	args := m.Called(pctx, req)
	return args.Error(0)
}

func (m *TodoRepoMock) InsertOneTodo(pctx context.Context, req *todo.Todo) (primitive.ObjectID, error) {
	args := m.Called(pctx, req)
	return args.Get(0).(primitive.ObjectID), args.Error(1)
}

func (m *TodoRepoMock) DeleteOneTodo(pctx context.Context, todoId string) (int64, error) {
	args := m.Called(pctx, todoId)
	return args.Get(0).(int64), args.Error(1)
}

func (m *TodoRepoMock) SearchTodo(pctx context.Context, text string) ([]*todo.Todo, error) {
	args := m.Called(pctx, text)
	return args.Get(0).([]*todo.Todo), args.Error(1)
}
