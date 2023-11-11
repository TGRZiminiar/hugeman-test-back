package server

import (
	todohandler "github.com/TGRZiminiar/hugeman-test-back/modules/todo/todoHandler"
	todorepository "github.com/TGRZiminiar/hugeman-test-back/modules/todo/todoRepository"
	todousecase "github.com/TGRZiminiar/hugeman-test-back/modules/todo/todoUsecase"
)

func (s *server) todoService() {
	repoTodo := todorepository.NewTodoRepository(s.db)
	usecaseTodo := todousecase.NewTodoUsecase(repoTodo)
	httpHandlerTodo := todohandler.NewTodoHttpHandler(s.cfg, usecaseTodo)

	todo := s.app.Group("/todo_v1")
	todo.GET("/list-todo", httpHandlerTodo.FindManyTodo)
	todo.GET("/todo/:todoId", httpHandlerTodo.FindOneTodo)
	todo.GET("/search-todo", httpHandlerTodo.SearchTodo)
	todo.POST("/todo", httpHandlerTodo.CreateTodo)
	todo.DELETE("/todo/:todoId", httpHandlerTodo.DeleteOneTodo)
	todo.PATCH("/todo", httpHandlerTodo.UpdateOneTodo)

}
