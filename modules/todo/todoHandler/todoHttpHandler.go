package todohandler

import (
	"context"
	"net/http"

	"github.com/TGRZiminiar/hugeman-test-back/config"
	"github.com/TGRZiminiar/hugeman-test-back/modules/todo"
	todousecase "github.com/TGRZiminiar/hugeman-test-back/modules/todo/todoUsecase"
	"github.com/TGRZiminiar/hugeman-test-back/pkg/request"
	"github.com/TGRZiminiar/hugeman-test-back/pkg/response"
	"github.com/gin-gonic/gin"
)

type (
	TodoHttpHandlerService interface {
		CreateItem(c *gin.Context)
		FindOneTodo(c *gin.Context)
		FindManyTodo(c *gin.Context)
		DeleteOneTodo(c *gin.Context)
	}

	todoHttpHandler struct {
		cfg         *config.Config
		todoUsecase todousecase.TodoUseCaseService
	}
)

func NewTodoHttpHandler(cfg *config.Config, todoUsecase todousecase.TodoUseCaseService) TodoHttpHandlerService {
	return &todoHttpHandler{
		cfg,
		todoUsecase,
	}
}

func (h *todoHttpHandler) CreateItem(c *gin.Context) {
	ctx := context.Background()

	wrapper := request.NewContextWrapper(c)

	req := new(todo.CreateTodoReq)

	if err := wrapper.Bind(req); err != nil {
		response.ErrResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	itemId, err := h.todoUsecase.InsertOneTodo(ctx, req)
	if err != nil {
		response.ErrResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	response.SuccessResponse(c, http.StatusCreated, itemId)
	return
}

func (h *todoHttpHandler) FindOneTodo(c *gin.Context) {
	ctx := context.Background()

	itemId := c.Param("todoId")

	if itemId == "" {
		response.ErrResponse(c, http.StatusBadRequest, "itemId is required")
		return
	}

	res, err := h.todoUsecase.FindOneTodo(ctx, itemId)
	if err != nil {
		response.ErrResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	response.SuccessResponse(c, http.StatusOK, res)
	return

}

func (h *todoHttpHandler) FindManyTodo(c *gin.Context) {
	ctx := context.Background()

	result, err := h.todoUsecase.FindManyTodo(ctx)
	if err != nil {
		response.ErrResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	response.SuccessResponse(c, http.StatusOK, result)
	return

}

func (h *todoHttpHandler) DeleteOneTodo(c *gin.Context) {
	ctx := context.Background()

	itemId := c.Param("todoId")

	if itemId == "" {
		response.ErrResponse(c, http.StatusBadRequest, "itemId is required")
		return
	}

	res, err := h.todoUsecase.DeleteOneTodo(ctx, itemId)
	if err != nil {
		response.ErrResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	response.SuccessResponse(c, http.StatusOK, res)
	return

}
