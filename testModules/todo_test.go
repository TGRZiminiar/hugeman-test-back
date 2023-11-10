package testmodules

import (
	"context"
	"errors"
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
	testFindOneObj struct {
		ctx      context.Context
		cfg      *config.Config
		isErr    bool
		expected *todo.TodoShowcase
		todoReq  string
	}
	testInsert struct {
		ctx      context.Context
		cfg      *config.Config
		isErr    bool
		expected *todo.TodoShowcase
		req      *todo.CreateTodoReq
	}
	testDel struct {
		ctx      context.Context
		cfg      *config.Config
		isErr    bool
		expected int64
		todoReq  string
	}
	testUpdate struct {
		ctx      context.Context
		cfg      *config.Config
		isErr    bool
		expected *todo.TodoShowcase
		todoReq  *todo.UpdateTodoReq
	}
)

func NewTestConfig() *config.Config {
	cfg := config.LoadConfig("../env/dev/.env.dev")
	return &cfg
}

func TestFindOne(t *testing.T) {
	repoMock := new(todorepository.TodoRepoMock)
	usecase := todousecase.NewTodoUsecase(repoMock)

	ctx := context.Background()
	cfg := NewTestConfig()

	resFindOneOK := primitive.NewObjectID()
	resFindOneFailed := primitive.NewObjectID()

	testsFindOne := []testFindOneObj{
		{
			ctx:     ctx,
			cfg:     cfg,
			todoReq: resFindOneOK.Hex(),
			isErr:   false,
			expected: &todo.TodoShowcase{
				Id:          resFindOneOK.Hex(),
				Title:       "Test1",
				Description: "Description1",
				CreatedAt:   time.Time{},
				UpdatedAt:   time.Time{},
				Image:       "data:image/png;base64,image1",
				Status:      "IN_PROGRESS",
			},
		},
		{
			ctx:      ctx,
			cfg:      cfg,
			todoReq:  resFindOneFailed.Hex(),
			isErr:    true,
			expected: nil,
		},
		{
			ctx:     ctx,
			cfg:     cfg,
			todoReq: resFindOneFailed.Hex(),
			isErr:   true,
			expected: &todo.TodoShowcase{
				Id:          resFindOneOK.Hex(),
				Title:       "Test3",
				Description: "Description3",
				CreatedAt:   time.Time{},
				UpdatedAt:   time.Time{},
				Image:       "data:image/png;base64,image3",
				Status:      "IN_PROGRESS",
			},
		},
	}

	repoMock.On("FindOneTodo", ctx, resFindOneOK.Hex()).Return(&todo.Todo{
		Id:          resFindOneOK,
		Title:       "Test1",
		Description: "Description1",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		Image:       "image1",
		Status:      "IN_PROGRESS",
	}, nil)

	repoMock.On("FindOneTodo", ctx, resFindOneFailed.Hex()).Return(&todo.Todo{}, errors.New("error: find one todo not found"))

	repoMock.On("FindOneTodo", ctx, resFindOneOK.Hex()).Return(&todo.Todo{
		Id:          resFindOneOK,
		Title:       "Test3",
		Description: "Description3",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		Image:       "data:image/png;base64,image3",
		Status:      "IN_PROGRESS",
	}, errors.New("error: find one todo not found"))

	fmt.Println("Find One Todo Test ->")
	for i, test := range testsFindOne {
		fmt.Printf("case -> %d\n", i+1)

		result, err := usecase.FindOneTodo(test.ctx, test.todoReq)

		if test.isErr {
			assert.NotEmpty(t, err)
		} else {
			result.CreatedAt = time.Time{}
			result.UpdatedAt = time.Time{}

			assert.Equal(t, test.expected, result)
		}
	}
}

func TestInsertOneTodo(t *testing.T) {
	repoMock := new(todorepository.TodoRepoMock)
	usecase := todousecase.NewTodoUsecase(repoMock)

	ctx := context.Background()
	cfg := NewTestConfig()

	resInsertOneOk := primitive.NewObjectID()
	resInsertOneFailed := primitive.NewObjectID()

	testsInsertOne := []testInsert{
		{
			ctx:   ctx,
			cfg:   cfg,
			isErr: false,
			req: &todo.CreateTodoReq{
				Title:       "Title1",
				Description: "Description1",
				Image:       "Image1",
				Status:      "IN_PROGRESS",
			},
			expected: &todo.TodoShowcase{
				Id:          resInsertOneOk.Hex(),
				Title:       "Title1",
				Description: "Description1",
				Image:       "data:image/png;base64,Image1",
				Status:      "IN_PROGRESS",
				CreatedAt:   time.Time{},
				UpdatedAt:   time.Time{},
			},
		},
		{
			ctx:   ctx,
			cfg:   cfg,
			isErr: true,
			req: &todo.CreateTodoReq{
				Title:       "Title2",
				Description: "Description2",
				Image:       "Image2",
				Status:      "IN_PROGRESS",
			},
			expected: nil,
		},
	}

	repoMock.On("FindOneTodo", ctx, resInsertOneOk.Hex()).Return(&todo.Todo{
		Id:          resInsertOneOk,
		Title:       "Title1",
		Description: "Description1",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		Image:       "Image1",
		Status:      "IN_PROGRESS",
	}, nil)

	repoMock.On("FindOneTodo", ctx, resInsertOneFailed.Hex()).Return(&todo.Todo{}, errors.New("error: find one todo not found"))

	repoMock.On("InsertOneTodo", ctx, &todo.Todo{
		Title:       "Title1",
		Description: "Description1",
		Image:       "Image1",
		Status:      "IN_PROGRESS",
	}).Return(resInsertOneOk, nil)

	repoMock.On("InsertOneTodo", ctx, &todo.Todo{
		Title:       "Title2",
		Description: "Description2",
		Image:       "Image2",
		Status:      "IN_PROGRESS",
	}).Return(resInsertOneFailed, errors.New("error: insert one todo failed"))

	fmt.Println("Insert One Todo Test ->")
	for i, test := range testsInsertOne {
		fmt.Printf("case -> %d\n", i+1)

		result, err := usecase.InsertOneTodo(test.ctx, test.req)
		if test.isErr {
			assert.NotEmpty(t, err)
		} else {
			result.CreatedAt = time.Time{}
			result.UpdatedAt = time.Time{}

			assert.Equal(t, test.expected, result)
		}
	}

}

func TestDeleteTodo(t *testing.T) {
	repoMock := new(todorepository.TodoRepoMock)
	usecase := todousecase.NewTodoUsecase(repoMock)

	ctx := context.Background()
	cfg := NewTestConfig()

	successDelId := primitive.NewObjectID()
	failedDelId := primitive.NewObjectID()

	testsInsertOne := []testDel{
		{
			ctx:      ctx,
			cfg:      cfg,
			isErr:    false,
			expected: int64(1),
			todoReq:  successDelId.Hex(),
		},
		{
			ctx:      ctx,
			cfg:      cfg,
			isErr:    true,
			expected: int64(-1),
			todoReq:  failedDelId.Hex(),
		},
	}

	repoMock.On("DeleteOneTodo", ctx, successDelId.Hex()).Return(int64(1), nil)

	repoMock.On("DeleteOneTodo", ctx, failedDelId.Hex()).Return(int64(-1), errors.New("error: deleteonetodo failed"))

	fmt.Println("Delete One Todo Test ->")
	for i, test := range testsInsertOne {
		fmt.Printf("case -> %d\n", i+1)

		result, err := usecase.DeleteOneTodo(test.ctx, test.todoReq)
		if test.isErr {
			assert.NotEmpty(t, err)
		} else {
			assert.Equal(t, test.expected, result)
		}
	}
}

func TestUpdateOneTodo(t *testing.T) {
	repoMock := new(todorepository.TodoRepoMock)
	usecase := todousecase.NewTodoUsecase(repoMock)

	ctx := context.Background()
	cfg := NewTestConfig()

	resUpdateOneOk := primitive.NewObjectID()
	resUpdateOneFailed := primitive.NewObjectID()

	testsUpdateOne := []testUpdate{
		{
			ctx:   ctx,
			cfg:   cfg,
			isErr: false,
			todoReq: &todo.UpdateTodoReq{
				Id:          resUpdateOneOk.Hex(),
				Title:       "Title1",
				Description: "Description1",
				Image:       "data:image/png;base64,Image1",
				Status:      "IN_PROGRESS",
			},
			expected: &todo.TodoShowcase{
				Id:          resUpdateOneOk.Hex(),
				Title:       "Title1",
				Description: "Description1",
				Image:       "data:image/png;base64,Image1",
				Status:      "IN_PROGRESS",
				CreatedAt:   time.Time{},
				UpdatedAt:   time.Time{},
			},
		},
		{
			ctx:   ctx,
			cfg:   cfg,
			isErr: true,
			todoReq: &todo.UpdateTodoReq{
				Id:          resUpdateOneFailed.Hex(),
				Title:       "Title2",
				Description: "",
				Image:       "data:image/png;base64,Image2",
				Status:      "IN_PROGRESS",
			},
			expected: &todo.TodoShowcase{
				Id:          resUpdateOneFailed.Hex(),
				Title:       "Title2",
				Description: "",
				Image:       "data:image/png;base64,Image2",
				Status:      "IN_PROGRESS",
				CreatedAt:   time.Time{},
				UpdatedAt:   time.Time{},
			},
		},
	}

	repoMock.On("FindOneTodo", ctx, resUpdateOneOk.Hex()).Return(&todo.Todo{
		Id:          resUpdateOneOk,
		Title:       "Title1",
		Description: "Description1",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		Image:       "Image1",
		Status:      "IN_PROGRESS",
	}, nil)

	repoMock.On("FindOneTodo", ctx, resUpdateOneFailed.Hex()).Return(&todo.Todo{}, errors.New("error: find one todo not found"))

	repoMock.On("UpdateOneTodo", ctx, &todo.TodoShowcase{
		Id:          resUpdateOneOk.Hex(),
		Title:       "Title1",
		Description: "Description1",
		Image:       "data:image/png;base64,Image1",
		Status:      "IN_PROGRESS",
	}).Return(nil)

	repoMock.On("UpdateOneTodo", ctx, &todo.TodoShowcase{
		Id:          resUpdateOneFailed.Hex(),
		Title:       "Title2",
		Description: "",
		Image:       "data:image/png;base64,Image2",
		Status:      "IN_PROGRESS",
	}).Return(errors.New("error: updateonetodo not found"))

	fmt.Println("Update One Todo Test ->")
	for i, test := range testsUpdateOne {
		fmt.Printf("case -> %d\n", i+1)

		result, err := usecase.UpdateOneTodo(test.ctx, test.todoReq)
		if test.isErr {
			assert.NotEmpty(t, err)
		} else {
			result.CreatedAt = time.Time{}
			result.UpdatedAt = time.Time{}

			assert.Equal(t, test.expected, result)
		}
	}

}
