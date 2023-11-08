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
	todo.POST("/todo", httpHandlerTodo.CreateItem)
	todo.DELETE("/todo/:todoId", httpHandlerTodo.DeleteOneTodo)
	todo.PATCH("/todo", httpHandlerTodo.UpdateOneTodo)

	// auth := s.app.Group("/auth_v1")

	// // HealthCheck

	// // auth.GET("", s.middleware.JwtAuthorization(s.middleware.RbacAuthorization(s.healthCheckService, []int{1, 0})))
	// auth.GET("", s.healthCheckService)
	// auth.POST("/auth/login", authHttpHandler.Login)
	// auth.POST("/auth/refresh-token", authHttpHandler.RefreshToken)
	// auth.POST("/auth/logout", authHttpHandler.Logout)
}
