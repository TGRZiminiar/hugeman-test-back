package testmodules

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/TGRZiminiar/hugeman-test-back/config"
	"github.com/TGRZiminiar/hugeman-test-back/modules/todo"
	todorepository "github.com/TGRZiminiar/hugeman-test-back/modules/todo/todoRepository"
	todousecase "github.com/TGRZiminiar/hugeman-test-back/modules/todo/todoUsecase"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	testLogin struct {
		ctx     context.Context
		cfg     *config.Config
		isErr   bool
		todoRes *todo.Todo
	}
)

func NewTestConfig() *config.Config {
	cfg := config.LoadConfig("../env/dev/.env.dev")
	return &cfg
}

func TestTodo(t *testing.T) {
	repoMock := new(todorepository.TodoRepoMock)
	usecase := todousecase.NewTodoUsecase(repoMock)

	ctx := context.Background()
	cfg := NewTestConfig()

	credentialIdSuccess := primitive.NewObjectID()
	credentialIdFailed := primitive.NewObjectID()

	tests := []testLogin{
		{
			ctx: ctx,
			cfg: cfg,
			todoRes: &todo.Todo{
				Id:          credentialIdSuccess,
				Title:       "Test1",
				Description: "Description1",
				CreatedAt:   time.Time{},
				UpdatedAt:   time.Time{},
				Image:       "image1",
				Status:      "IN_PROGESS",
			},
		},
		{
			ctx: ctx,
			cfg: cfg,
		},
		{
			ctx: ctx,
			cfg: cfg,
		},
	}

	repoMock.On("FindOneTodo", ctx, credentialIdSuccess.Hex()).Return(&todo.Todo{
		Id:          credentialIdSuccess,
		Title:       "Test1",
		Description: "Description1",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		Image:       "image1",
		Status:      "IN_PROGESS",
	}, nil)

	repoMock.On("FindOneTodo", ctx, credentialIdSuccess.Hex()).Return(&todo.Todo{
		Id:          credentialIdSuccess,
		Title:       "Test1",
		Description: "Description1",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		Image:       "image1",
		Status:      "IN_PROGESS",
	}, nil)

	for i, test := range tests {
		fmt.Printf("case -> %d\n", i+1)

		result, err := usecase.Login(test.ctx, test.cfg, test.req)

		if test.isErr {
			assert.NotEmpty(t, err)
		} else {
			result.CreatedAt = time.Time{}
			result.UpdatedAt = time.Time{}
			result.Credential.CreatedAt = time.Time{}
			result.Credential.UpdatedAt = time.Time{}

			assert.Equal(t, test.expected, result)
		}
	}
}
