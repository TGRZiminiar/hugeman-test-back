package todohandler

import (
	"context"
	"encoding/base64"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/TGRZiminiar/hugeman-test-back/config"
	"github.com/TGRZiminiar/hugeman-test-back/modules/todo"
	todousecase "github.com/TGRZiminiar/hugeman-test-back/modules/todo/todoUsecase"
	"github.com/TGRZiminiar/hugeman-test-back/pkg/request"
	"github.com/TGRZiminiar/hugeman-test-back/pkg/response"
	"github.com/gin-gonic/gin"
)

type (
	TodoHttpHandlerService interface {
		CreateTodo(c *gin.Context)
		FindOneTodo(c *gin.Context)
		FindManyTodo(c *gin.Context)
		DeleteOneTodo(c *gin.Context)
		UpdateOneTodo(c *gin.Context)
		SearchTodo(c *gin.Context)
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

func (h *todoHttpHandler) CreateTodo(c *gin.Context) {

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

	todoId := c.Param("todoId")

	if todoId == "" {
		response.ErrResponse(c, http.StatusBadRequest, "todoId is required")
		return
	}

	res, err := h.todoUsecase.FindOneTodo(ctx, todoId)
	if err != nil {
		response.ErrResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	response.SuccessResponse(c, http.StatusOK, res)

}
func (h *todoHttpHandler) FindManyTodo(c *gin.Context) {
	ctx := context.Background()
	var limit, page int
	var sort string = "date"

	limitQ, ok := c.GetQuery("limit")
	if !ok {
		limit = 5
	} else {
		limit, _ = strconv.Atoi(limitQ)
	}
	pageQ, ok := c.GetQuery("page")
	if !ok {
		page = 1
	} else {
		page, _ = strconv.Atoi(pageQ)
	}
	sortQ := c.Query("sort")
	switch sortQ {
	case "created_at", "status", "title":
		sort = sortQ
	default:
		sort = "created_at"
	}

	result, err := h.todoUsecase.FindManyTodo(ctx, page, limit, sort)
	if err != nil {
		response.ErrResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	response.SuccessResponse(c, http.StatusOK, result)
}

func (h *todoHttpHandler) DeleteOneTodo(c *gin.Context) {
	ctx := context.Background()

	todoId := c.Param("todoId")

	if todoId == "" {
		response.ErrResponse(c, http.StatusBadRequest, "todoId is required")
		return
	}

	res, err := h.todoUsecase.DeleteOneTodo(ctx, todoId)
	if err != nil {
		response.ErrResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	response.SuccessResponse(c, http.StatusOK, gin.H{
		"msg":   "Delete todo success",
		"count": res,
	})

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
		if strings.Contains(req.Image, "data:image/png;base64,") {
			result := strings.Replace(req.Image, "data:image/png;base64,", "", -1)
			req.Image = result
		} else {
			req.Image = ""
		}
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
	response.SuccessResponse(c, http.StatusOK, todos)
}

func (h *todoHttpHandler) SearchTodo(c *gin.Context) {
	ctx := context.Background()

	search, ok := c.GetQuery("search")
	if !ok {
		response.ErrResponse(c, http.StatusBadRequest, "search query is required")
		return
	}

	res, err := h.todoUsecase.SearchTodo(ctx, search)
	if err != nil {
		response.ErrResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	response.SuccessResponse(c, http.StatusOK, res)

}
