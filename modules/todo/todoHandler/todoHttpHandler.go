package todohandler

import (
	"context"
	"encoding/base64"
	"io"
	"log"
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
		UpdateOneTodo(c *gin.Context)
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

	image, err := c.FormFile("image")
	req := new(todo.CreateTodoReq)
	if image != nil {
		if err != nil {
			log.Printf("Error: Validate File Failed: %s", err.Error())
			response.ErrResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		file, err := image.Open()
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		data, err := io.ReadAll(file)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		req.Image = base64.StdEncoding.EncodeToString(data)
	} else {
		req.Image = ""
	}

	if err := wrapper.Bind(req); err != nil {
		log.Printf("Error: Validate Request Failed: %s", err.Error())
		response.ErrResponse(c, http.StatusBadRequest, "error: some field is required or invalid")
		return
	}

	todo, err := h.todoUsecase.InsertOneTodo(ctx, req)
	if err != nil {
		log.Printf("Error: Insert Todo Failed: %s", err.Error())
		response.ErrResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	response.SuccessResponse(c, http.StatusCreated, todo)
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

}

func (h *todoHttpHandler) FindManyTodo(c *gin.Context) {
	ctx := context.Background()

	result, err := h.todoUsecase.FindManyTodo(ctx)
	if err != nil {
		response.ErrResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	response.SuccessResponse(c, http.StatusOK, result)

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

}

func (h *todoHttpHandler) UpdateOneTodo(c *gin.Context) {
	ctx := context.Background()

	wrapper := request.NewContextWrapper(c)

	image, err := c.FormFile("image")
	req := new(todo.UpdateTodoReq)
	if image != nil {
		if err != nil {
			log.Printf("Error: Validate File Failed: %s", err.Error())
			response.ErrResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		file, err := image.Open()
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		data, err := io.ReadAll(file)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		req.Image = base64.StdEncoding.EncodeToString(data)
	} else {
		req.Image = ""
	}

	if err := wrapper.Bind(req); err != nil {
		log.Printf("Error: Validate Request Failed: %s", err.Error())
		response.ErrResponse(c, http.StatusBadRequest, "error: some field is required or invalid")
		return
	}

	todos, err := h.todoUsecase.UpdateOneTodo(ctx, req)
	if err != nil {
		log.Printf("Error: Update Todo Failed: %s", err.Error())
		response.ErrResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	response.SuccessResponse(c, http.StatusCreated, todos)
}
